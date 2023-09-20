package controller

import (
	"context"

	"github.com/chuxin0816/Scaffold/models"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func SignUpHandler(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	p := new(models.ParamSignUp)
	if err := ctx.BindAndValidate(p); err != nil {
		hlog.Error("SignUp with invalid param, err:%v", err)
		ctx.JSON(consts.StatusOK, utils.H{"msg": "error"})
		return
	}

	// 调用service层处理业务逻辑
	// service.SignUp(p)
	ctx.JSON(consts.StatusOK, utils.H{"msg": "success"})
}
