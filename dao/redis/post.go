package redis

import (
	"context"
	"time"

	"github.com/chuxin0816/bluebell/models"
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

func GetPostIDsInOrder(ppl *models.ParamPostList) ([]string, error) {
	// 从redis中获取postID列表
	key := getRedisKey(KeyPostTimeZSet)
	if ppl.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (ppl.Page - 1) * ppl.Size
	stop := start + ppl.Size - 1
	// 按分数从大到小查询
	return rdb.ZRevRange(context.Background(), key, start, stop).Result()
}

func GetPostVoteData(ids []string) (voteData []int64, err error) {
	voteData = make([]int64, 0, len(ids))
	// 提前查好每个post的投票数
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(context.Background(), key, "1", "1")
	}
	cmder, err := pipeline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	for _, cmd := range cmder {
		v := cmd.(*redis.IntCmd).Val()
		voteData = append(voteData, v)
	}
	return
}
