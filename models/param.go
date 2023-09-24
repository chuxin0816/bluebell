package models

type ParamRegister struct {
	Username        string `json:"username" vd:"len($)>0&&len($)<20"`
	Password        string `json:"password" vd:"len($)>7&&len($)<20&&$==(ConfirmPassword)$"`
	ConfirmPassword string `json:"confirm_password" vd:"len($)>7&&len($)<20"`
}

type ParamLogin struct {
	Username string `json:"username" vd:"len($)>0&&len($)<20"`
	Password string `json:"password" vd:"len($)>7&&len($)<20"`
}
