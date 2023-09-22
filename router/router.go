package router

import (
	"context"
	"fmt"

	"github.com/chuxin0816/Scaffold/config"
	"github.com/chuxin0816/Scaffold/controller"
	"github.com/chuxin0816/Scaffold/middleware"
	"github.com/chuxin0816/Scaffold/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func SetUp(conf *config.HertzConfig) *server.Hertz {
	h := server.Default(server.WithHostPorts(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	))

	h.GET("/ping", middleware.AuthMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		response.Success(ctx, utils.H{"user": ctx.MustGet("user")}, "pong")
	})
	h.POST("/register", controller.RegisterHandler)
	h.POST("/login", controller.LoginHandler)

	return h
}
