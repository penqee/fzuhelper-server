// Code generated by Kitex v0.9.1. DO NOT EDIT.

package templateservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	template "github.com/west2-online/fzuhelper-server/kitex_gen/template"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Ping": kitex.NewMethodInfo(
		pingHandler,
		newTemplateServicePingArgs,
		newTemplateServicePingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	templateServiceServiceInfo                = NewServiceInfo()
	templateServiceServiceInfoForClient       = NewServiceInfoForClient()
	templateServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return templateServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return templateServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return templateServiceServiceInfoForClient
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
	serviceName := "TemplateService"
	handlerType := (*template.TemplateService)(nil)
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
		"PackageName": "template",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*template.TemplateServicePingArgs)
	realResult := result.(*template.TemplateServicePingResult)
	success, err := handler.(template.TemplateService).Ping(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTemplateServicePingArgs() interface{} {
	return template.NewTemplateServicePingArgs()
}

func newTemplateServicePingResult() interface{} {
	return template.NewTemplateServicePingResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, req *template.PingRequest) (r *template.PingResponse, err error) {
	var _args template.TemplateServicePingArgs
	_args.Req = req
	var _result template.TemplateServicePingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
