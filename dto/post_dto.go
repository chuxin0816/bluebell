package dto

import (
	"time"

	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/models"
)

type PostDto struct {
	PostID     int64         `json:"post_id"`
	Status     int           `json:"status"`
	Community  *CommunityDto `json:"community"`
	AuthorName string        `json:"author_name"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func ToPostDto(post *models.Post) (*PostDto, error) {
	author, exist := mysql.CheckUserIDExist(post.AuthorID)
	if !exist {
		return nil, mysql.ErrorUserNotExist
	}
	authorName := author.Username
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		return nil, mysql.ErrorCommunityNotFound
	}
	return &PostDto{
		PostID:     post.PostID,
		Status:     post.Status,
		AuthorName: authorName,
		Community:  ToCommunityDto(community),
		Title:      post.Title,
		Content:    post.Content,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}, nil
}
