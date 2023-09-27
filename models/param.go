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
	CommunityID int    `json:"community_id,string" vd:"$>0"`
	Status      int    `json:"status,string" vd:"$>0"`
	Title       string `json:"title" vd:"len($)>0"`
	Content     string `json:"content" vd:"len($)>0"`
}

type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" vd:"$>0"`
	Direction int8  `json:"direction,string" vd:"$>-2&&$<2"`
}
