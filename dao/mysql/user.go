package mysql

import (
	"github.com/chuxin0816/Scaffold/models"
)

func QueryByUserName(username string) (err error) {
	user := models.User{Username: username}
	err = db.Where("username=?", username).First(&user).Error
	return
}

func InsertUser(){
	
}
