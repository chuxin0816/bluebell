package redis

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePreVote    = 432
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
	ErrorVoteRepeated   = errors.New("不允许重复投票")
)

func VotePost(userID, postID string, val float64) error {
	// 判断帖子时间是否超出限制
	postTime := rdb.ZScore(context.Background(), getRedisKey(KeyPostTimeZSet), postID).Val()
	if (float64(time.Now().Unix()) - postTime) > oneWeekInSecond {
		return ErrorVoteTimeExpire
	}
	// 查看当前用户给当前帖子的投票记录
	ov := rdb.ZScore(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	if ov == val {
		return ErrorVoteRepeated
	}
	// 计算分数差值
	var op float64
	if val > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(val - ov)
	// 开启事务
	pipeline := rdb.TxPipeline()
	// 更新帖子分数
	pipeline.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID).Err()
	// 记录用户投票记录
	if val == 0 {
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), userID).Err()
	}
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
		Score:  val,
		Member: userID,
	}).Err()
	_, err := pipeline.Exec(context.Background())
	return err
}
