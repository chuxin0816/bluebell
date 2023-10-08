package router

import (
	"context"
	"fmt"
	"time"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/controller"

	// _ "github.com/chuxin0816/bluebell/docs"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	// "github.com/hertz-contrib/swagger"
	// swaggerFiles "github.com/swaggo/files"
)

func SetUp(conf *config.HertzConfig) *server.Hertz {
	h := server.Default(server.WithHostPorts(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	))

	h.Use(cors.Default())

	// h.LoadHTMLFiles("./template/index.html")
	// h.Static("/static", ".")
	// h.GET("/", func(c context.Context, ctx *app.RequestContext) {
	// ctx.HTML(consts.StatusOK, "index.html", nil)
	// })

	// pprof.Register(h)
	// h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	h.Use(middleware.RatelimitMiddleware(time.Millisecond, 1000))

	v1 := h.Group("/api/v1")

	v1.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})
	v1.POST("/signup", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)

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
