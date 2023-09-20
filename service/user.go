package service

import (
	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/models"
	"github.com/chuxin0816/Scaffold/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	// 查询用户是否存在
	err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	// 生成UID
	userID := snowflake.GenerateID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存用户信息
	err = mysql.InsertUser(user)
	return err
}
