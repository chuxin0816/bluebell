package models

import "time"

type User struct {
	ID        uint   `gorm:"type:bigint(20)"`
	Gender    int    `gorm:"type:tinyint(4);default:0"`
	UserID    int64  `gorm:"type:bigint(20);not null;unique"`
	Username  string `gorm:"type:varchar(64);not null;unique"`
	Password  string `gorm:"type:varchar(64);not null"`
	Email     string `gorm:"type:varchar(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
