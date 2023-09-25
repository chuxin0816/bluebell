package controller

import (
	"context"
	"strconv"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dto"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type IPostController interface {
	Create(c context.Context, ctx *app.RequestContext)
	Show(c context.Context, ctx *app.RequestContext)
	List(c context.Context, ctx *app.RequestContext)
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

func (pp *PostController) Show(c context.Context, ctx *app.RequestContext) {
	// 从请求获取参数
	postID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		hlog.Error("Get post with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	// 调用service层处理业务逻辑
	post, err := service.GetPost(postID)
	if err != nil {
		hlog.Error("Get post with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	// 转换为dto对象
	postDto, err := dto.ToPostDto(post)
	if err != nil {
		hlog.Error("Get post with dto error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, utils.H{"post": postDto}, "")
}

func (pp *PostController) List(c context.Context, ctx *app.RequestContext) {
	// 从请求获取参数
	pageNum, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		pageSize = 5
	}
	// 调用service层处理业务逻辑
	postList, err := service.GetPostList(pageNum, pageSize)
	if err != nil {
		hlog.Error("Get post list with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	postDtoList, err := dto.ToPostDtoList(postList)
	if err != nil {
		hlog.Error("Get post list with dto error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, utils.H{"post_list": postDtoList}, "")
}
