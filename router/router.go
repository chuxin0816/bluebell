package router

import (
	"context"
	"fmt"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/controller"
	"github.com/chuxin0816/bluebell/dto"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func SetUp(conf *config.HertzConfig) *server.Hertz {
	h := server.Default(server.WithHostPorts(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	))

	h.GET("/ping", middleware.AuthMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		user := ctx.MustGet(response.CtxUserKey).(*models.User)
		response.Success(ctx, utils.H{"user": dto.ToUserDto(user)}, "pong")
	})
	h.POST("/register", controller.RegisterHandler)
	h.POST("/login", controller.LoginHandler)

	return h
}
