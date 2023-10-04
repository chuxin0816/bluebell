package controller

import (
	"encoding/json"
	"testing"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/response"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func init() {
	if err := mysql.Init(&config.MysqlConfig{
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DBName:   "bluebell",
	}); err != nil {
		panic(err)
	}
}

func TestList(t *testing.T) {
	h := server.Default()
	communityController, err := NewCommunityController()
	if err != nil {
		t.Error("NewCommunityController with error: ", err)
	}
	h.GET("/list", communityController.List)
	w := ut.PerformRequest(h.Engine, "GET", "/list", nil,
		ut.Header{
			Key:   "Connection",
			Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, `{"code":1000,"msg":"请求成功","data":{"community_list":[{"community_id":"1","community_name":"GO","introduction":"Go语言，由Google开源，拥有活跃的社区","created_time":"2023-09-23T20:20:20+08:00"},{"community_id":"2","community_name":"Python","introduction":"Python语言，是目前排名第一的语言，语法简单易学","created_time":"2020-09-09T11:11:23+08:00"}],"total":2}}`, string(resp.Body()))
}

func TestShow(t *testing.T) {
	h := server.Default()
	communityController, err := NewCommunityController()
	if err != nil {
		t.Error("NewCommunityController with error: ", err)
	}
	h.GET("/show/:id", communityController.Show)
	w := ut.PerformRequest(h.Engine, "GET", "/show/1", nil, ut.Header{
		Key:   "Connection",
		Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, `{"code":1000,"msg":"请求成功","data":{"community":{"community_id":"1","community_name":"GO","introduction":"Go语言，由Google开源，拥有活跃的社区","created_time":"2023-09-23T20:20:20+08:00"}}}`, string(resp.Body()))
}

func TestShow2(t *testing.T) { // 不存在的社区
	h := server.Default()
	communityController, err := NewCommunityController()
	if err != nil {
		t.Error("NewCommunityController with error: ", err)
	}
	h.GET("/show/:id", communityController.Show)
	w := ut.PerformRequest(h.Engine, "GET", "/show/10086", nil, ut.Header{
		Key:   "Connection",
		Value: "close"})
	resp := w.Result()
	responseData := new(response.ResponseData)
	err = json.Unmarshal(resp.Body(), responseData)
	if err != nil {
		t.Error("json.Unmarshal with error: ", err)
	}
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "服务器繁忙", responseData.Msg)
}
