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
	"errors"
	"io"
	"time"

	dkio "github.com/CodapeWild/devkit/io"
	networkv1 "github.com/CodapeWild/opentracing-agent/network/v1"
	"github.com/google/martian/v3/log"
)

type TraceService struct {
	networkv1.UnimplementedTracesReportServiceServer
}

func (ts *TraceService) SendTrace(src networkv1.TracesReportService_SendTraceServer) error {
	for {
		trace, err := src.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				src.SendAndClose(grpc_ok.With(CommonRespWithMessage("EOF from client side")).Response)

				return nil
			}
			log.Debugf(err.Error())
			continue
		}

		// todo: send trace to io
	}
}

func (ts *TraceService) SendSpan(src networkv1.TracesReportService_SendSpanServer) error {
	spanCollector := dkio.NewCacheAndFlush(100, 5*time.Second)
	spanCollector.Subscribe(func(data any) {
		set, ok := data.([]any)
		if !ok {
			log.Debugf("type assertion failed want []any get %T", data)

			return
		}

		var trace = &networkv1.Trace{Trace: make([]*networkv1.Span, len(set))}
		for i := range set {
			span, ok := set[i].(*networkv1.Span)
			if !ok {
				log.Debugf("type assertion failed want span get %T", set[i])
				break
			}
			trace.Trace[i] = span
		}

		// todo: send trace to io
	})
	spanCollector.Start()
	defer spanCollector.Close()

	for {
		span, err := src.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				src.SendAndClose(grpc_error.With(CommonRespWithMessage("EOF from client side")).Response)

				return nil
			}
			log.Debugf(err.Error())
			continue
		}

		spanCollector.Publish(span)
	}
}
