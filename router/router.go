package router

import (
	"fmt"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/controller"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func SetUp(conf *config.HertzConfig) *server.Hertz {
	h := server.Default(server.WithHostPorts(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	))

	h.POST("/register", controller.RegisterHandler)
	h.POST("/login", controller.LoginHandler)
	h.GET("/info", middleware.AuthMiddleware(), controller.InfoHandler)

	communityRouter := h.Group("/community", middleware.AuthMiddleware())
	{
		communityController := controller.NewCommunityController()
		communityRouter.GET("/", communityController.List)
	}

	return h
}
