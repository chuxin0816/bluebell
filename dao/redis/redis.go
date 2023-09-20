package redis

import (
	"fmt"

	"github.com/chuxin0816/Scaffold/config"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(conf *config.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
