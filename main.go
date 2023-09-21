package main

import (
	"fmt"

	"github.com/chuxin0816/Scaffold/config"
	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/dao/redis"
	"github.com/chuxin0816/Scaffold/logger"
	"github.com/chuxin0816/Scaffold/pkg/snowflake"
	"github.com/chuxin0816/Scaffold/router"
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
