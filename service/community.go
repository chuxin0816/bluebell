package service

import (
	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/models"
)

func GetCommunityList() (communities []models.Community, err error) {
	communityList, err := mysql.GetCommunityList()
	return communityList, err
}

func GetCommunityByID(communityID string) (community *models.Community, err error) {
	community, err = mysql.GetCommunityByID(communityID)
	return community, err
}
