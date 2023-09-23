package controller

import (
	"context"
	"errors"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dto"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func RegisterHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamSignUp)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	// 调用service层处理业务逻辑
	token, err := service.Register(p)
	if err != nil {
		hlog.Error("SignUp with service error: ", err)
		if errors.Is(err, mysql.ErrorUserExist) {
			response.Error(ctx, response.CodeUserExist, "")
			return
		}
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, utils.H{"token": token}, "注册成功")
}

func LoginHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamLogin)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	// 调用service层处理业务逻辑
	token, err := service.Login(p)
	if err != nil {
		hlog.Error("Login with service error: ", err)
		response.Error(ctx, response.CodeInvalidPassword, "")
		return
	}
	response.Success(ctx, utils.H{"token": token}, "登陆成功")
}

func InfoHandler(c context.Context, ctx *app.RequestContext) {
	user := ctx.MustGet("user").(*models.User)
	response.Success(ctx, utils.H{"user": dto.ToUserDto(user)}, "")
}
