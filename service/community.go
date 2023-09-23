package service

import (
	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/models"
)

func GetCommunityList() (communities []models.Community, err error) {
	communityList, err := mysql.GetCommunityList()
	return communityList, err
}
