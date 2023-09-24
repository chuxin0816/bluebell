package dto

import (
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type PostDto struct {
	PostID      int64  `gorm:"type:bigint(20);not null"`
	AuthorID    int64  `gorm:"type:bigint(20)"`
	CommunityID int    `gorm:"type:bigint(20)"`
	Status      int    `gorm:"type:tinyint(4);default:1"`
	Title       string `gorm:"type:varchar(128)"`
	Content     string `gorm:"type:varchar(8192)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToPostDto(post *models.Post) *PostDto {
	return &PostDto{
		PostID:      post.PostID,
		AuthorID:    post.AuthorID,
		CommunityID: post.CommunityID,
		Title:       post.Title,
		Content:     post.Content,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
