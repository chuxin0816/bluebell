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

func GetPostList(ppl *models.ParamPostList) (postList []*models.Post, err error) {
	// 从redis中获取postID列表
	ids, err := redis.GetPostIDsInOrder(ppl)
	if err != nil {
		return nil, err
	}
	// 提前查好每个post的投票数
	voteData, err := redis.GetPostVoteData(ids)
	// 根据postID列表从mysql中获取post列表
	postList, err = mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	for idx, post := range postList {
		post.VoteNum = voteData[idx]
	}
	return
}

func VotePost(userID int64, pv *models.ParamVoteData) error {
	return redis.VotePost(strconv.FormatInt(userID, 10), strconv.FormatInt(pv.PostID, 10), float64(pv.Direction))
}
