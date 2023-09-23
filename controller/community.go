package controller

import (
	"context"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/response"
	"github.com/chuxin0816/bluebell/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type ICommunityController interface {
	List(c context.Context, ctx *app.RequestContext)
}

type CommunityController struct{}

func NewCommunityController() ICommunityController {
	mysql.NewCommunity()
	return &CommunityController{}
}

func (community *CommunityController) List(c context.Context, ctx *app.RequestContext) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		hlog.Error("GetCommunityList with mysql error: ", err)
		response.Error(ctx, response.CodeServerBusy, "")
	}
	total := len(communityList)
	response.Success(ctx, utils.H{"community_list": communityList, "total": total}, "")
}
