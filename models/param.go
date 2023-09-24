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

type ParamPost struct {
	PostID      int64  `json:"id" vd:"$>0"`
	AuthorID    int64  `json:"author_id" vd:"$>0"`
	CommunityID int    `json:"community_id" vd:"$>0"`
	Status      int    `json:"status" vd:"required"`
	Title       string `json:"title" vd:"len($)>0"`
	Content     string `json:"content" vd:"len($)>0"`
}
