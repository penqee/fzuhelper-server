// Code generated by hertz generator.

package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/west2-online/fzuhelper-server/biz/model/api"
)

// GetScores .
// @router /academic/scores [GET]
func GetScores(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ScoresRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ScoresResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetGPA .
// @router /academic/gpa [GET]
func GetGPA(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GPARequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GPAResp)

	c.JSON(consts.StatusOK, resp)
}

// GetCredit .
// @router /academic/credit [GET]
func GetCredit(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreditRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreditResp)

	c.JSON(consts.StatusOK, resp)
}

// GetUnifiedExam .
// @router /academic/unifiedExam [GET]
func GetUnifiedExam(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UnifiedExamRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UnifiedExamResp)

	c.JSON(consts.StatusOK, resp)
}

// GetPlan .
// @router /academic/plan [GET]
func GetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PlanRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PlanResp)

	c.JSON(consts.StatusOK, resp)
}
