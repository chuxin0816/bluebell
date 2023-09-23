package controller

import (
	"context"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type communityController struct{}

func NewCommunityController() *communityController {
	mysql.NewCommunity()
	return &communityController{}
}

func (community *communityController) GetCommunityList(c context.Context, ctx *app.RequestContext) {
	communityList := mysql.GetCommunityList()
	response.Success(ctx, utils.H{"community_list": communityList}, "")
}
