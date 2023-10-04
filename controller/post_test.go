package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/dao/redis"
	"github.com/chuxin0816/bluebell/middleware"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

const token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTE2ODE1OTE2Nzk3OTUyLCJleHAiOjE2OTY0MzE2NjcsImlzcyI6ImNodXhpbiJ9.LwfvleB-pwydQQ4iN9hrOMhi5cBgfjGAQtL9JMXwXUM"

func init() {
	if err := redis.Init(&config.RedisConfig{Host: "127.0.0.1", Port: 6379, Password: "", DB: 0}); err != nil {
		fmt.Printf("redis init failed, err:%v\n", err)
		return
	}
}
func TestCreate(t *testing.T) {
	h := server.Default()
	postController, err := NewPostController()
	if err != nil {
		panic(err)
	}
	h.POST("/create", middleware.AuthMiddleware(), postController.Create)
	body := `{
		"community_id":"1",
		"status":"1",
		"title":"test post",
		"content":"It is just a test post"
	}`
	w := ut.PerformRequest(h.Engine, "POST", "/create", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body)},
		ut.Header{
			Key:   "Authorization",
			Value: token,
		},
		ut.Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err = json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("Unmarshal response data with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "创建成功", responseData.Msg)
}

func TestShowPost(t *testing.T) {
	h := server.Default()
	postController, err := NewPostController()
	if err != nil {
		panic(err)
	}
	h.GET("/show/:id", postController.Show)
	w := ut.PerformRequest(h.Engine, "GET", "/show/3482196870434816", nil, ut.Header{
		Key:   "Connection",
		Value: "close",
	})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, `{"code":1000,"msg":"请求成功","data":{"post":{"post_id":"3482196870434816","status":"1","community":{"community_id":"2","community_name":"Python","introduction":"Python语言，是目前排名第一的语言，语法简单易学","created_time":"2020-09-09T11:11:23+08:00"},"author_name":"小碗","vote_num":"0","title":"Python简单教程","content":"首先打开b站，然后搜索Python，随便打开一个视频开学！","created_time":"2023-09-27T22:37:00.431+08:00"}}}`, string(resp.Body()))
}

func TestPostList(t *testing.T) {
	h := server.Default()
	postController, err := NewPostController()
	if err != nil {
		panic(err)
	}
	h.GET("/posts", postController.List)
	w := ut.PerformRequest(h.Engine, "GET", "/posts", nil, ut.Header{
		Key:   "Connection",
		Value: "close",
	})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err = json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("Unmarshal response data with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "请求成功", responseData.Msg)
}

func TestVote(t *testing.T) {
	h := server.Default()
	postController, err := NewPostController()
	if err != nil {
		panic(err)
	}
	body := `{
		"post_id":"3482196870434816",
		"direction":"1"
	}`
	h.POST("/vote", middleware.AuthMiddleware(), postController.Vote)
	w := ut.PerformRequest(h.Engine, "POST", "/vote", &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body)},
		ut.Header{
			Key:   "Authorization",
			Value: token,
		},
		ut.Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
		ut.Header{
			Key:   "Connection",
			Value: "close",
		})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err = json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("Unmarshal response data with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "服务器繁忙", responseData.Msg)
}
