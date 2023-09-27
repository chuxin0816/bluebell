package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreatePost(postID string) error {
	// 开启事务
	pipeline := rdb.TxPipeline()
	// 帖子发帖时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子初始分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec(context.Background())
	return err
}
