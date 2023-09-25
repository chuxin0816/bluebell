package dto

import (
	"strconv"
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type UserDto struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	CreateAt time.Time
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		UserID:   strconv.FormatInt(user.UserID, 10),
		Username: user.Username,
		Email:    user.Email,
		Gender:   strconv.Itoa(user.Gender),
		CreateAt: user.CreatedAt,
	}
}
