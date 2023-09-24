package mysql

import (
	"errors"

	"github.com/chuxin0816/bluebell/models"
)

var (
	ErrorPostNotFound    = errors.New("没有找到相关帖子")
	ErrorAutoMigratePost = errors.New("迁移帖子表失败")
)

func NewPost() error {
	return db.AutoMigrate(&models.Post{})
}

func CreatePost(p *models.Post) error {
	return db.Create(p).Error
}
