package controller

import (
	"context"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app"
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
	communityList := mysql.GetCommunityList()
	response.Success(ctx, utils.H{"community_list": communityList}, "")
}
