package sequence

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

type Redis struct {
	rdb *redis.Client
	key string // Redis 键名，默认 "sequence:global"
}

// NewRedis 创建发号器
func NewRedis(addr, password string, db int) Sequence {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &Redis{
		rdb: rdb,
		key: "sequence:global",
	}
}

// Next 获取下一个全局递增 ID
func (r *Redis) Next() (seq uint64, err error) {
	// INCR 原子自增，返回新值
	val, err := r.rdb.Incr(r.key).Result()
	if err != nil {
		logx.Errorw("redis.Incr failed",
			logx.LogField{Key: "err", Value: err.Error()},
			logx.LogField{Key: "key", Value: r.key})
		return 0, err
	}
	return uint64(val), nil
}
