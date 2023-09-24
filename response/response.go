package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Response(ctx *app.RequestContext, httpStatus int, code ResCode, data utils.H, msg string) {
	ctx.JSON(httpStatus, utils.H{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func Success(ctx *app.RequestContext, data utils.H, msg string) {
	if msg == "" {
		msg = CodeSuccess.Message()
	}
	Response(ctx, consts.StatusOK, CodeSuccess, data, msg)
}

func Error(ctx *app.RequestContext, code ResCode, msg string) {
	if msg == "" {
		msg = code.Message()
	}
	Response(ctx, consts.StatusOK, code, nil, msg)
}
