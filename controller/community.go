package controller

import (
	"context"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dto"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type ICommunityController interface {
	List(c context.Context, ctx *app.RequestContext)
	Show(c context.Context, ctx *app.RequestContext)
}

type CommunityController struct{}

func NewCommunityController() ICommunityController {
	mysql.NewCommunity()
	return &CommunityController{}
}

func (cc *CommunityController) List(c context.Context, ctx *app.RequestContext) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		hlog.Error("GetCommunityList with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
	}
	total := len(communityList)
	response.Success(ctx, utils.H{"community_list": dto.ToCommunityDtoList(communityList), "total": total}, "")
}

func (cc *CommunityController) Show(c context.Context, ctx *app.RequestContext) {
	communityID := ctx.Param("id")
	community, err := service.GetCommunityByID(communityID)
	if err != nil {
		hlog.Error("GetCommunityByID with service error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
	}
	response.Success(ctx, utils.H{"community": dto.ToCommunityDto(community)}, "")
}
