package mysql

import (
	"errors"

	"github.com/chuxin0816/bluebell/models"
)

var (
	ErrorCommunityNotFound    = errors.New("没有找到相关社区")
	ErrorAutoMigrateCommunity = errors.New("迁移社区表失败")
)

func NewCommunity() error {
	err := db.AutoMigrate(&models.Community{})
	if err != nil {
		return ErrorAutoMigrateCommunity
	}
	return nil
}

func GetCommunityList() (communities []models.Community, err error) {
	db.Find(&communities)
	if len(communities) == 0 {
		return nil, ErrorCommunityNotFound
	}
	return communities, nil
}

func GetCommunityByID(communityID string) (community *models.Community, err error) {
	community = new(models.Community)
	db.Where("community_id = ?", communityID).First(community)
	if community.ID == 0 {
		return nil, ErrorCommunityNotFound
	}
	return community, nil
}
