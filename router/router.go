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
	v1 := h.Group("/api/v1")
	v1.POST("/signup", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/info", middleware.AuthMiddleware(), controller.InfoHandler)

	communityRouter := v1.Group("/community", middleware.AuthMiddleware())
	{
		communityController := controller.NewCommunityController()
		communityRouter.GET("/", communityController.List)
		communityRouter.GET("/:id", communityController.Show)
	}

	return h
}
