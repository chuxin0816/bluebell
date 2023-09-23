package mysql

import (
	"errors"

	"github.com/chuxin0816/bluebell/models"
)

var (
	ErrorCommunityNotFound = errors.New("没有找到相关社区")
)

func NewCommunity() {
	db.AutoMigrate(&models.Community{})
}

func GetCommunityList() (communities []models.Community, err error) {
	db.Find(&communities)
	if len(communities) == 0 {
		return nil, ErrorCommunityNotFound
	}
	return communities, nil
}
