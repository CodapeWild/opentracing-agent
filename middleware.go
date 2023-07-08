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
	"net/http"

	dkhttp "github.com/CodapeWild/devkit/net/http"
	networkv1 "github.com/CodapeWild/opentracing-agent/network/v1"
)

// For HTTP status
var (
	http200 = dkhttp.NewJSONRespMessage(http.StatusOK, dkhttp.JSONRespMsgWithVersion(v1))
	http400 = dkhttp.NewJSONRespMessage(http.StatusBadRequest, dkhttp.JSONRespMsgWithVersion(v1))
	http500 = dkhttp.NewJSONRespMessage(http.StatusInternalServerError, dkhttp.JSONRespMsgWithVersion(v1))
)

func checkHeadersFailed(resp http.ResponseWriter, req *http.Request) {
	dkhttp.NewJSONRespMessage(http.StatusBadRequest, dkhttp.JSONRespMsgWithMessage("invalid HTTP request headers")).WriteBy(resp)
}

// For gRPC status
type CommonRespOption func(cresp *CommonResponse)

func CommonRespWithStatus(status int) CommonRespOption {
	return func(cresp *CommonResponse) { cresp.Status = int32(status) }
}

func CommonRespWithVersion(version string) CommonRespOption {
	return func(cresp *CommonResponse) { cresp.Version = version }
}

func CommonRespWithMessage(msg string) CommonRespOption {
	return func(cresp *CommonResponse) { cresp.Message = msg }
}

func CommonRespWithPayload(coding string, payload []byte) CommonRespOption {
	return func(cresp *CommonResponse) {
		cresp.Coding = coding
		cresp.Payload = payload
	}
}

type CommonResponse struct {
	*networkv1.Response
}

func (cresp *CommonResponse) With(opts ...CommonRespOption) *CommonResponse {
	for _, opt := range opts {
		opt(cresp)
	}

	return cresp
}

func NewCommonResponse(opts ...CommonRespOption) *CommonResponse {
	cresp := &CommonResponse{&networkv1.Response{}}
	for _, opt := range opts {
		opt(cresp)
	}

	return cresp
}

const (
	gRPCStatusOK    = 1000
	gRPCStatusError = 2000
)

var (
	grpc_ok    = NewCommonResponse(CommonRespWithStatus(gRPCStatusOK), CommonRespWithVersion(v1))
	grpc_error = NewCommonResponse(CommonRespWithStatus(gRPCStatusError), CommonRespWithVersion(v1))
)
