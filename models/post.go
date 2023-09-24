package models

import "time"

type Post struct {
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
