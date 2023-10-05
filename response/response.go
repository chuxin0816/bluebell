package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"`
}

func Response(ctx *app.RequestContext, httpStatus int, code ResCode, data any, msg string) {
	ctx.JSON(httpStatus, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(ctx *app.RequestContext, data any, msg string) {
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
