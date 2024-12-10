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

	"github.com/cloudwego/hertz/pkg/app"

	api "github.com/west2-online/fzuhelper-server/api/model/api"
	"github.com/west2-online/fzuhelper-server/api/pack"
	"github.com/west2-online/fzuhelper-server/api/rpc"
	"github.com/west2-online/fzuhelper-server/kitex_gen/common"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
)

// GetTermsList .
// @router /api/v1/terms/list [GET]
func GetTermsList(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(common.TermListRequest)
	res, err := rpc.GetTermsList(ctx, req)
	if err != nil {
		pack.RespError(c, err)
		return
	}

	resp := new(api.TermListResponse)
	resp.TermLists = pack.BuildTermList(res)

	pack.RespData(c, resp.TermLists)
}

// GetTerm .
// @router /api/v1/terms/info [GET]
func GetTerm(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.TermRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	res, err := rpc.GetTerm(ctx, &common.TermRequest{
		Term: req.Term,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	resp := new(api.TermResponse)
	resp.TermInfo = pack.BuildTermInfo(res)

	pack.RespData(c, resp.TermInfo)
}
