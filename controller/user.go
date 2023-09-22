package controller

import (
	"context"
	"errors"

	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/models"
	"github.com/chuxin0816/Scaffold/response"
	"github.com/chuxin0816/Scaffold/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RegisterHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamSignUp)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam)
		return
	}
	// 调用service层处理业务逻辑
	if err := service.Register(p); err != nil {
		hlog.Error("SignUp with service error: ", err)
		if errors.Is(err, mysql.ErrorUserExist) {
			response.Error(ctx, response.CodeUserExist)
			return
		}
		response.Error(ctx, response.CodeServerBusy)
		return
	}
	response.Success(ctx, nil)
}

func LoginHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamLogin)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam)
		return
	}
	// 调用service层处理业务逻辑
	if err := service.Login(p); err != nil {
		hlog.Error("Login with service error: ", err)
		response.Error(ctx, response.CodeInvalidPassword)
		return
	}
	response.Success(ctx, nil)
}
