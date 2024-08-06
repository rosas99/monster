package main

import (
	"context"
	"github.com/redis/go-redis/v9"

	"github.com/zeromicro/go-zero/core/stringx"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

const (
	randomLen       = 16
	tolerance       = 500
	millisPerSecond = 1000
)

var (
	// 这里使用 go-redis 的 Lua 脚本功能
	lockScript = redis.NewScript(`
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
			return "OK"
		else
			return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
		end
	`)
	delScript = redis.NewScript(`
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`)
)

// RedisLock 是使用 Redis 实现的分布式锁。
type RedisLock struct {
	store   *redis.Client
	seconds uint32
	key     string
	id      string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRedisLock 返回一个新的 RedisLock 实例。
func NewRedisLock(store *redis.Client, key string) *RedisLock {
	return &RedisLock{
		store:   store,
		key:     key,
		id:      stringx.Randn(randomLen), // 假设 stringx.Randn 是可用的
		seconds: 30,                       // 默认超时时间
	}
}

// AcquireCtx 尝试获取锁，使用提供的上下文。
func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	val, err := lockScript.Run(ctx, rl.store, []string{rl.key}, rl.id, int64(seconds*millisPerSecond+tolerance)).Result()
	if err != nil {
		log.Printf("Error on acquiring lock for %s: %v", rl.key, err)
		return false, err
	}
	return val == "OK", nil
}

// ReleaseCtx 尝试释放锁，使用提供的上下文。
func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	val, err := delScript.Run(ctx, rl.store, []string{rl.key}, rl.id).Result()
	if err != nil {
		return false, err
	}
	return val.(int64) == 1, nil
}

// SetExpire 设置锁的超时时间（秒）。
func (rl *RedisLock) SetExpire(seconds uint32) {
	atomic.StoreUint32(&rl.seconds, seconds)
}
