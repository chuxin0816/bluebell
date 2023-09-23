package redis

import (
	"context"
	"fmt"

	"github.com/chuxin0816/bluebell/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(conf *config.RedisConfig) error {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	err := rdb.Ping(ctx).Err()
	return err
}

func Close() {
	_ = rdb.Close()
}
