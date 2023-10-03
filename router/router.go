package router

import (
	"fmt"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/controller"
	_ "github.com/chuxin0816/bluebell/docs"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func SetUp(conf *config.HertzConfig) *server.Hertz {
	h := server.Default(server.WithHostPorts(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	))
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	v1 := h.Group("/api/v1")
	v1.POST("/signup", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/info", middleware.AuthMiddleware(), controller.InfoHandler)

	communityRouter := v1.Group("/community")
	{
		communityController, err := controller.NewCommunityController()
		if err != nil {
			hlog.Error("NewCommunityController with error: ", err)
		}
		communityRouter.GET("/", communityController.List)
		communityRouter.GET("/:id", communityController.Show)
	}

	postRouter := v1.Group("/")
	{
		postController, err := controller.NewPostController()
		if err != nil {
			hlog.Error("NewPostController with err: ", err)
		}
		postRouter.GET("/post/:id", postController.Show)
		postRouter.GET("/posts2", postController.List)
		postRouter.Use(middleware.AuthMiddleware())
		postRouter.POST("/post", postController.Create)
		postRouter.POST("/vote", postController.Vote)
	}

	return h
}
