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
	"errors"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/west2-online/fzuhelper-server/api/model/api"
	"github.com/west2-online/fzuhelper-server/api/model/model"
	"github.com/west2-online/fzuhelper-server/api/pack"
	"github.com/west2-online/fzuhelper-server/api/rpc"
	"github.com/west2-online/fzuhelper-server/kitex_gen/paper"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
)

// ListDirFiles .
// @router /api/v1/paper/list [GET]
func ListDirFiles(ctx context.Context, c *app.RequestContext) {
	var err error

	path := c.DefaultQuery("path", "")
	if path == "" {
		pack.RespError(c, errno.ParamError.WithError(errors.New("path is empty")))
		return
	}

	res, err := rpc.GetDirFilesRPC(ctx, &paper.ListDirFilesRequest{
		Path: path,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	resp := new(api.ListDirFilesResponse)
	resp.Dir = pack.BuildUpYunFileDir(res)
	pack.RespData(c, resp.Dir)
}

// GetDownloadUrl .
// @router /api/v1/paper/download [GET]
func GetDownloadUrl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetDownloadUrlRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamError.WithError(err))
		return
	}

	url, err := rpc.GetDownloadUrlRPC(ctx, &paper.GetDownloadUrlRequest{
		Filepath: req.Filepath,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}

	resp := new(api.GetDownloadUrlResponse)
	resp.URL = url

	pack.RespData(c, resp)
}

// ListDirFilesForAndroid .
// @Description 历年卷兼容，查询文件目录
// @router /api/v1/list [GET]
func ListDirFilesForAndroid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ListDirFilesForAndroidRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespErrorInPaper(c, errno.ParamError.WithError(err))
		return
	}

	if req.GetPath() == "" {
		pack.RespErrorInPaper(c, errno.ParamError.WithError(errors.New("path is empty")))
		return
	}

	res, err := rpc.GetDirFilesRPC(ctx, &paper.ListDirFilesRequest{
		Path: req.GetPath(),
	})
	if err != nil {
		pack.RespErrorInPaper(c, err)
		return
	}

	data := &model.PaperData{
		BasePath: res.BasePath,
		Files:    res.Files,
		Folders:  res.Folders,
	}

	pack.RespDataInPaper(c, data)
}

// GetDownloadUrlForAndroid .
// @Description 历年卷兼容，获取url以供下载
// @router /api/v1/downloadUrl [GET]
func GetDownloadUrlForAndroid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetDownloadUrlForAndroidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespErrorInPaper(c, errno.ParamError.WithError(err))
		return
	}

	url, err := rpc.GetDownloadUrlRPC(ctx, &paper.GetDownloadUrlRequest{
		Filepath: req.Filepath,
	})
	if err != nil {
		pack.RespErrorInPaper(c, err)
		return
	}

	data := &model.PaperUrlData{
		URL: url,
	}

	pack.RespDataInPaper(c, data)
}
