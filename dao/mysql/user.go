package mysql

import (
	"errors"

	"github.com/chuxin0816/Scaffold/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	var user models.User
	err = db.Where("username = ?", username).First(&user).Error
	if user.ID > 0 {
		return errors.New("用户已存在")
	}
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
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
