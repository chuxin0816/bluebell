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
)

type IPostController interface {
	Create(c context.Context, ctx *app.RequestContext)
	Show(c context.Context, ctx *app.RequestContext)
	List(c context.Context, ctx *app.RequestContext)
	Vote(c context.Context, ctx *app.RequestContext)
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
	// 调用service层处理业务逻辑
	err = service.CreatePost(pp, userID)
	if err != nil {
		hlog.Error("Create post with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, nil, "创建成功")
}

func (pc *PostController) Show(c context.Context, ctx *app.RequestContext) {
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
	response.Success(ctx, postDto, "")
}

// List 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} []dto.PostDto
// @Router /posts2 [get]
func (pc *PostController) List(c context.Context, ctx *app.RequestContext) {
	// 从请求获取参数
	ppl := &models.ParamPostList{
		Page:        1,
		Size:        5,
		Order:       models.OrderTime,
		CommunityID: 0,
	}
	err := ctx.BindAndValidate(ppl)
	if err != nil {
		hlog.Error("Get post list with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	// 调用service层处理业务逻辑
	postList, err := service.GetPostList(ppl)
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
	response.Success(ctx, postDtoList, "")
}

func (pc *PostController) Vote(c context.Context, ctx *app.RequestContext) {
	// 从请求获取参数
	pv := new(models.ParamVoteData)
	err := ctx.BindAndValidate(pv)
	if err != nil {
		hlog.Error("Vote post with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	userID := ctx.MustGet(middleware.CtxUserKey).(*models.User).UserID
	// 调用service层处理业务逻辑
	err = service.VotePost(userID, pv)
	if err != nil {
		hlog.Error("Vote post with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, nil, "投票成功")
}
