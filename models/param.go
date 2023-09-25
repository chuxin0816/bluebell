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
	PostID      int64  `json:"id"`
	AuthorID    int64  `json:"author_id"`
	CommunityID int    `json:"community_id" vd:"$>0"`
	Status      int    `json:"status" vd:"$>0"`
	Title       string `json:"title" vd:"len($)>0"`
	Content     string `json:"content" vd:"len($)>0"`
}
