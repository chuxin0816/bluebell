package dto

import (
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type PostDto struct {
	ID          int64  `gorm:"type:bigint(20)"`
	PostID      int64  `gorm:"type:bigint(20);not null"`
	Title       string `gorm:"type:varchar(128)"`
	Content     string `gorm:"type:varchar(8192)"`
	AuthorID    int64  `gorm:"type:bigint(20)"`
	CommunityID int64  `gorm:"type:bigint(20)"`
	Status      uint   `gorm:"type:tinyint(4);default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToPostDto(post *models.Post) *PostDto {
	return &PostDto{
		PostID:      post.PostID,
		Title:       post.Title,
		Content:     post.Content,
		AuthorID:    post.AuthorID,
		CommunityID: post.CommunityID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
