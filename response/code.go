package response

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
}

func (code ResCode) Message() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
