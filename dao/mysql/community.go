package mysql

import "github.com/chuxin0816/bluebell/models"

func NewCommunity() {
	db.AutoMigrate(&models.Community{})
}

func GetCommunityList() (communities []models.Community) {
	db.Find(&communities)
	return communities
}
