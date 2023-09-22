package service

import (
	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/models"
	"github.com/chuxin0816/Scaffold/pkg/jwt"
	"github.com/chuxin0816/Scaffold/pkg/snowflake"
)

func Register(p *models.ParamSignUp) error {
	// 查询用户是否存在
	if _, exist := mysql.CheckUsernameExist(p.Username); exist {
		return mysql.ErrorUserExist
	}
	// 生成UID
	userID := snowflake.GenerateID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存用户信息
	err := mysql.InsertUser(user)
	return err
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return "", err
	}
	// 发放token

	return jwt.GenToken(user.UserID)

}
