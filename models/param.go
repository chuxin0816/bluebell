package models

type ParamSignUp struct {
	Username   string `json:"username" vd:"len($)>0&&len($)<20"`
	Password   string `json:"password" vd:"len($)>7&&len($)<20&&$==(RePassword)$"`
	RePassword string `json:"re_password" vd:"len($)>7&&len($)<20"`
	Email      string `json:"email" vd:"len($)==0||email($)"`
}
