/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.11.3. DO NOT EDIT.

package commonservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	common "github.com/west2-online/fzuhelper-server/kitex_gen/common"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetTermsList": kitex.NewMethodInfo(
		getTermsListHandler,
		newCommonServiceGetTermsListArgs,
		newCommonServiceGetTermsListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetTerm": kitex.NewMethodInfo(
		getTermHandler,
		newCommonServiceGetTermArgs,
		newCommonServiceGetTermResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	commonServiceServiceInfo                = NewServiceInfo()
	commonServiceServiceInfoForClient       = NewServiceInfoForClient()
	commonServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commonServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commonServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commonServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CommonService"
	handlerType := (*common.CommonService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "common",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func getTermsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonServiceGetTermsListArgs)
	realResult := result.(*common.CommonServiceGetTermsListResult)
	success, err := handler.(common.CommonService).GetTermsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonServiceGetTermsListArgs() interface{} {
	return common.NewCommonServiceGetTermsListArgs()
}

func newCommonServiceGetTermsListResult() interface{} {
	return common.NewCommonServiceGetTermsListResult()
}

func getTermHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonServiceGetTermArgs)
	realResult := result.(*common.CommonServiceGetTermResult)
	success, err := handler.(common.CommonService).GetTerm(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonServiceGetTermArgs() interface{} {
	return common.NewCommonServiceGetTermArgs()
}

func newCommonServiceGetTermResult() interface{} {
	return common.NewCommonServiceGetTermResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetTermsList(ctx context.Context, req *common.TermListRequest) (r *common.TermListResponse, err error) {
	var _args common.CommonServiceGetTermsListArgs
	_args.Req = req
	var _result common.CommonServiceGetTermsListResult
	if err = p.c.Call(ctx, "GetTermsList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTerm(ctx context.Context, req *common.TermRequest) (r *common.TermResponse, err error) {
	var _args common.CommonServiceGetTermArgs
	_args.Req = req
	var _result common.CommonServiceGetTermResult
	if err = p.c.Call(ctx, "GetTerm", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
