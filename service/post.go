package service

import (
	"strconv"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dao/redis"
	"github.com/chuxin0816/bluebell/models"
	"github.com/chuxin0816/bluebell/pkg/snowflake"
)

func CreatePost(pp *models.ParamPost, userID int64) (err error) {
	_, err = mysql.GetCommunityByID(pp.CommunityID)
	if err != nil {
		return mysql.ErrorCommunityNotFound
	}
	// 生成postID
	postID := snowflake.GenerateID()
	p := &models.Post{
		PostID:      postID,
		AuthorID:    userID,
		CommunityID: pp.CommunityID,
		Status:      pp.Status,
		Title:       pp.Title,
		Content:     pp.Content,
	}
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	return redis.CreatePost(strconv.FormatInt(postID, 10))
}

func GetPost(postID int64) (post *models.Post, err error) {
	return mysql.GetPost(postID)
}

func GetPostList(pageNum, pageSize int) (postList []*models.Post, err error) {
	return mysql.GetPostList(pageNum, pageSize)
}

func VotePost(userID int64, pv *models.ParamVoteData) error {
	return redis.VotePost(strconv.FormatInt(userID, 10), strconv.FormatInt(pv.PostID, 10), float64(pv.Direction))
}
