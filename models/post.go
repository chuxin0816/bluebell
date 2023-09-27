package models

import "time"

type Post struct {
	ID          int64  `gorm:"type:bigint(20)"`
	PostID      int64  `gorm:"type:bigint(20);not null"`
	AuthorID    int64  `gorm:"type:bigint(20)"`
	VoteNum     int64  `gorm:"type:bigint(20);default:0"`
	CommunityID int    `gorm:"type:bigint(10)"`
	Status      int    `gorm:"type:tinyint(4);default:1"`
	Title       string `gorm:"type:varchar(128)"`
	Content     string `gorm:"type:varchar(8192)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
