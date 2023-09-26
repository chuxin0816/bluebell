package service

import (
	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/pkg/snowflake"
)

func CreatePost(pp *models.ParamPost) (err error) {
	_, err = mysql.GetCommunityByID(pp.CommunityID)
	if err != nil {
		return mysql.ErrorCommunityNotFound
	}
	// 生成postID
	postID := snowflake.GenerateID()
	p := &models.Post{
		PostID:      postID,
		AuthorID:    pp.AuthorID,
		CommunityID: pp.CommunityID,
		Status:      pp.Status,
		Title:       pp.Title,
		Content:     pp.Content,
	}
	return mysql.CreatePost(p)
}

func GetPost(postID int64) (post *models.Post, err error) {
	return mysql.GetPost(postID)
}

func GetPostList(pageNum, pageSize int) (postList []*models.Post, err error) {
	return mysql.GetPostList(pageNum, pageSize)
}
