package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/chuxin0816/bluebell/models"
	"github.com/redis/go-redis/v9"
)

func getIDs(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	stop := start + size - 1
	// 按分数从大到小查询
	return rdb.ZRevRange(context.Background(), key, start, stop).Result()
}

func CreatePost(postID, communityID string) error {
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
	cKey := getRedisKey(KeyCommunitySetPF + communityID)
	pipeline.SAdd(context.Background(), cKey, postID)
	_, err := pipeline.Exec(context.Background())
	return err
}

func GetPostIDsInOrder(ppl *models.ParamPostList) ([]string, error) {
	// 从redis中获取postID列表
	key := getRedisKey(KeyPostTimeZSet)
	if ppl.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDs(key, ppl.Page, ppl.Size)
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

func GetCommunityPostIDsInOrder(ppl *models.ParamPostList) ([]string, error) {

	orderKey := getRedisKey(KeyPostTimeZSet)
	if ppl.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(ppl.CommunityID))
	key := orderKey + ":" + strconv.Itoa(ppl.CommunityID)
	// 利用缓存减少ZInterStore次数
	if rdb.Exists(context.Background(), key).Val() < 1 {
		// 不存在，需计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Keys:      []string{cKey, orderKey},
			Aggregate: "MAX",
		})
		pipeline.Expire(context.Background(), key, time.Second*60)
		_, err := pipeline.Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	return getIDs(key, ppl.Page, ppl.Size)
}
