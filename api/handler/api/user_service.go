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

// Code generated by hertz generator.

package api

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/west2-online/fzuhelper-server/api/model/api"
	"github.com/west2-online/fzuhelper-server/api/pack"
	"github.com/west2-online/fzuhelper-server/api/rpc"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
	"github.com/west2-online/fzuhelper-server/pkg/constants"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"github.com/west2-online/fzuhelper-server/pkg/logger"
)

// GetLoginData .
// @router /api/v1/user/login [GET]
func GetLoginData(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetLoginDataRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.GetLoginData: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}
	resp := new(api.GetLoginDataResponse)
	id, cookies, err := rpc.GetLoginDataRPC(ctx, &user.GetLoginDataRequest{
		Id:       req.ID,
		Password: req.Password,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	resp.ID = id
	resp.Cookies = cookies
	pack.RespData(c, resp)
}

// ValidateCode .
// @router /api/v1/jwch/user/validateCode [POST]
func ValidateCode(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ValidateCodeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.ValidateCode: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	request := new(protocol.Request)
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(constants.ValidateCodeURL)
	request.SetFormData(
		map[string]string{
			"image": req.Image,
		},
	)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.String(http.StatusOK, res.BodyBuffer().String())
}

// ValidateCodeForAndroid .
// @router /api/login/validateCode [POST]
func ValidateCodeForAndroid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ValidateCodeForAndroidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		logger.Errorf("api.ValidateCodeForAndroid: BindAndValidate error %v", err)
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	request := new(protocol.Request)
	request.SetMethod(consts.MethodPost)
	request.SetRequestURI(constants.ValidateCodeURL)
	request.SetFormData(
		map[string]string{
			"image": req.ValidateCode,
		},
	)

	res := new(protocol.Response)

	if err = ClientSet.HzClient.Do(ctx, request, res); err != nil {
		pack.RespError(c, err)
		return
	}

	c.String(http.StatusOK, res.BodyBuffer().String())
}
