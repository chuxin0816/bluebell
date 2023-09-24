package dto

import (
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type UserDto struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   uint   `json:"gender"`
	CreateAt time.Time
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Gender:   user.Gender,
		CreateAt: user.CreatedAt,
	}
}
