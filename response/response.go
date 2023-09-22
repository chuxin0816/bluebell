package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Response(ctx *app.RequestContext, httpStatus int, code ResCode, msg string, data utils.H) {
	ctx.JSON(httpStatus, utils.H{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func Success(ctx *app.RequestContext, data utils.H) {
	Response(ctx, consts.StatusOK, CodeSuccess, CodeSuccess.Message(), data)
}

func Error(ctx *app.RequestContext, code ResCode) {
	Response(ctx, consts.StatusOK, code, code.Message(), nil)
}

func ErrorWithMsg(ctx *app.RequestContext, code ResCode, msg string) {
	Response(ctx, consts.StatusOK, code, msg, nil)
}
