package controller

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/chuxin0816/bluebell/pkg/snowflake"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func init() {
	if err := snowflake.Init("2023-09-18", 1); err != nil {
		panic(err)
	}
}
func TestRegisterHandler(t *testing.T) {
	h := server.Default()
	h.POST("/register", RegisterHandler)
	body := `{
		"username":"小碗",
		"password":"12345678",
		"confirm_password":"12345678"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/register", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err := json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "用户名已存在", responseData.Msg)
}

func TestRegisterHandler2(t *testing.T) { // 两次密码不一致
	h := server.Default()
	h.POST("/register", RegisterHandler)
	body := `{
		"username":"小碗",
		"password":"12345678",
		"confirm_password":"12345679"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/register", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err := json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "请求参数错误", responseData.Msg)
}

func TestLogin(t *testing.T) {
	h := server.Default()
	h.POST("/login", LoginHandler)
	body := `{
		"username":"小碗",
		"password":"12345678"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/login", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err := json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "登陆成功", responseData.Msg)
}

func TestLogin2(t *testing.T) { // 用户名不存在
	h := server.Default()
	h.POST("/login", LoginHandler)
	body := `{
		"username":"小碗碗",
		"password":"12345678"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/login", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err := json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "用户名或密码错误", responseData.Msg)
}

func TestLogin3(t *testing.T) { // 密码错误
	h := server.Default()
	h.POST("/login", LoginHandler)
	body := `{
		"username":"小碗",
		"password":"12345679"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/login", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err := json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "用户名或密码错误", responseData.Msg)
}
