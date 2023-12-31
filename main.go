package main

// @title bluebell API
// @version 1.0
// @description bluebell API文档
// @license.name Apache 2.0
// @contact.name chuxin
// @contact.url github.com/chuxin0816
// @host 127.0.0.1:8888
// @BasePath /api/v1
import (
	"fmt"

	"github.com/chuxin0816/bluebell/config"
	"github.com/chuxin0816/bluebell/dao/mysql"
	"github.com/chuxin0816/bluebell/dao/redis"
	"github.com/chuxin0816/bluebell/logger"
	"github.com/chuxin0816/bluebell/pkg/snowflake"
	"github.com/chuxin0816/bluebell/router"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("config init failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("logger init failed, err:%v\n", err)
		return
	}
	// 初始化mysql
	if err := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init failed, err:%v\n", err)
		return
	}
	// 初始化redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	// 初始化雪花算法
	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("snowflake init failed, err:%v\n", err)
		return
	}
	// 注册路由
	h := router.SetUp(config.Conf.HertzConfig)
	// 启动服务
	h.Spin()
}
