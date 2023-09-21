package controller

import (
	"context"

	"github.com/chuxin0816/Scaffold/models"
	"github.com/chuxin0816/Scaffold/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamSignUp)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		ctx.JSON(consts.StatusOK, utils.H{"message": "注册失败"})
		return
	}
	// 调用service层处理业务逻辑
	if err := service.Register(p); err != nil {
		hlog.Error("SignUp with service error: ", err)
		ctx.JSON(consts.StatusOK, utils.H{"message": "注册失败"})
		return
	}
	ctx.JSON(consts.StatusOK, utils.H{"message": "注册成功"})
}

func LoginHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamLogin)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param: ", err)
		ctx.JSON(consts.StatusOK, utils.H{"message": "用户名或密码错误"})
		return
	}
	// 调用service层处理业务逻辑
	if err := service.Login(p); err != nil {
		hlog.Error("Login with service error: ", err)
		ctx.JSON(consts.StatusOK, utils.H{"message": "用户名或密码错误"})
		return
	}
	ctx.JSON(consts.StatusOK, utils.H{"message": "登录成功"})
}
