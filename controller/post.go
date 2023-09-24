package controller

import (
	"context"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type IPostController interface {
	Create(c context.Context, ctx *app.RequestContext)
}

type PostController struct{}

func NewPostController() (IPostController, error) {
	err := mysql.NewPost()
	return &PostController{}, err
}

func (pc *PostController) Create(c context.Context, ctx *app.RequestContext) {
	// 从请求中获取参数
	pp := new(models.ParamPost)
	err := ctx.BindAndValidate(pp)
	if err != nil {
		hlog.Error("Create post with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	userID := ctx.MustGet(middleware.CtxUserKey).(*models.User).UserID
	pp.AuthorID = userID
	// 调用service层处理业务逻辑
	err = service.CreatePost(pp)
	if err != nil {
		hlog.Error("Create post with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, nil, "创建成功")
}
