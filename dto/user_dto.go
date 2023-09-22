package dto

import "github.com/chuxin0816/bluebell/models"

type UserDto struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		UserID:   user.UserID,
		Username: user.Username,
	}
}
