package service

import (
	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/models"
	"github.com/chuxin0816/Scaffold/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 查询用户是否存在
	mysql.QueryByUserName(p.Username)
	// 生成UID
	snowflake.GenerateID()
	// 保存用户信息
	mysql.InsertUser()
}
