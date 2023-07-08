/*
 *   Copyright (c) 2023 CodapeWild
 *   All rights reserved.

 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at

 *   http://www.apache.org/licenses/LICENSE-2.0

 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/CodapeWild/devkit/bufpool"
	dkhttp "github.com/CodapeWild/devkit/net/http"
	networkv1 "github.com/CodapeWild/opentracing-agent/network/v1"
	"github.com/google/martian/v3/log"
	"google.golang.org/protobuf/proto"
)

var pattern = map[string]http.HandlerFunc{
	"/v1/trace": dkhttp.CheckHeaders(handleTrace(v1), checkHeadersFailed, validHeaders),
	"/v1/span":  dkhttp.CheckHeaders(handleSpan(v1), checkHeadersFailed, validHeaders),
}

func handleTrace(version string) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		bufpool.MakeUseOfBuffer(func(buf *bytes.Buffer) {
			io.Copy(buf, req.Body)
			trace, err := decodeTraceWithVersion(version, buf)
			if err != nil {
				log.Debugf(err.Error())
				http400.With(dkhttp.JSONRespMsgWithMessage(err.Error())).WriteBy(resp)

				return
			}
			// todo: send trace to native datacenter

			http200.WriteBy(resp)
		})
	}
}

func handleSpan(version string) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		bufpool.MakeUseOfBuffer(func(buf *bytes.Buffer) {
			io.Copy(buf, req.Body)
			span, err := decodeSpanWithVersion(version, buf)
			if err != nil {
				log.Debugf(err.Error())
				http400.With(dkhttp.JSONRespMsgWithMessage(err.Error())).WriteBy(resp)

				return
			}
			// todo: send span to native datacenter

			http200.WriteBy(resp)
		})
	}
}

func decodeTraceWithVersion(version string, buf *bytes.Buffer) (*networkv1.Trace, error) {
	switch version {
	case v1:
		var trace = &networkv1.Trace{}

		return trace, proto.Unmarshal(buf.Bytes(), trace)
	default:
		return nil, fmt.Errorf("unrecognized version: %s", version)
	}
}

func decodeSpanWithVersion(version string, buf *bytes.Buffer) (*networkv1.Span, error) {
	switch version {
	case v1:
		var span = &networkv1.Span{}

		return span, proto.Unmarshal(buf.Bytes(), span)
	default:
		return nil, fmt.Errorf("unrecognized version: %s", version)
	}
}
