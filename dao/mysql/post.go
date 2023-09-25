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

func GetPost(postID int64) (post *models.Post, err error) {
	post = new(models.Post)
	db.Where("post_id=?", postID).First(post)
	if post.ID == 0 {
		return nil, ErrorPostNotFound
	}
	return post, nil
}

func GetPostList(pageNum, pageSize int) (postList []models.Post, err error) {
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&postList)
	if len(postList) == 0 {
		return nil, ErrorPostNotFound
	}
	return postList, nil
}
