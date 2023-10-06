package controller

import (
	"context"
	"strconv"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dto"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ICommunityController interface {
	List(c context.Context, ctx *app.RequestContext)
	Show(c context.Context, ctx *app.RequestContext)
}

type CommunityController struct{}

func NewCommunityController() (ICommunityController, error) {
	err := mysql.NewCommunity()
	return &CommunityController{}, err
}

func (cc *CommunityController) List(c context.Context, ctx *app.RequestContext) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		hlog.Error("GetCommunityList with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, dto.ToCommunityDtoList(communityList), "")
}

func (cc *CommunityController) Show(c context.Context, ctx *app.RequestContext) {
	// 获取参数
	communityID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		hlog.Error("GetCommunityByID with invalid param: ", err)
		response.Error(ctx, response.CodeInvalidParam, "")
		return
	}
	// 调用service层处理业务逻辑
	community, err := service.GetCommunityByID(communityID)
	if err != nil {
		hlog.Error("GetCommunityByID with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
		return
	}
	response.Success(ctx, dto.ToCommunityDto(community), "")
}
