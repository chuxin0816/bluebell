package models

import "time"

type User struct {
	ID        uint   `gorm:"type:bigint(20)"`
	UserID    uint   `gorm:"type:bigint(20);not null;unique"`
	Username  string `gorm:"type:varchar(64);not null;unique"`
	Password  string `gorm:"type:varchar(64);not null"`
	Email     string `gorm:"type:varchar(64);not null"`
	Gender    uint   `gorm:"type:tinyint(4);default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
