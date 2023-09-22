package mysql

import (
	"errors"

	"github.com/chuxin0816/Scaffold/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorUserExist       = errors.New("用户名已存在")
	ErrorUserNotExist    = errors.New("用户名不存在")
	ErrorBcrypt          = errors.New("加密密码失败")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查指定用户ID的用户是否存在
func CheckUserIDExist(userID int64) (user *models.User, exist bool) {
	user = new(models.User)
	db.Where("user_id = ?", userID).First(user)
	return user, user.ID != 0
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUsernameExist(username string) (user *models.User, exist bool) {
	user = new(models.User)
	db.Where("username = ?", username).First(&user)
	return user, user.ID != 0
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) error {
	// 加密密码
	password, err := bcrypt.GenerateFromPassword([]byte((user.Password)), bcrypt.DefaultCost)
	if err != nil {
		return ErrorBcrypt
	}
	user.Password = string(password)

	// 创建用户
	err = db.Create(user).Error
	return err
}

func Login(user *models.User) error {
	password := user.Password
	db.Where("username = ?", user.Username).First(&user)
	if user.ID == 0 {
		return ErrorUserNotExist
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return ErrorInvalidPassword
	}
	return nil
}
