package mysql

import (
	"errors"
	"fmt"

	"github.com/chuxin0816/Scaffold/models"
	"golang.org/x/crypto/bcrypt"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (exist bool) {
	var user models.User
	db.Where("username = ?", username).First(&user)
	return user.ID != 0
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) error {
	// 加密密码
	password, err := bcrypt.GenerateFromPassword([]byte((user.Password)), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("加密密码失败")
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
		return errors.New("用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New(fmt.Sprint("用户:", user.Username, " 密码错误"))
	}
	return nil
}
