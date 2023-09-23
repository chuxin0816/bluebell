package models

import "time"

type Community struct {
	ID            int    `gorm:"type:int(11)"`
	CommunityID   int    `gorm:"type:int(10);not null;unique"`
	CommunityName string `gorm:"varchar(128);not null;unique"`
	Introduction  string `gorm:"varchar(256);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
