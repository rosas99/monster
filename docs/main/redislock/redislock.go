package redislock

import (
	"context"
	rands "crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/pkg/log"
	"io"
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

type RedisLock struct {
	store   *redis.Client
	seconds uint32
	key     string
	id      string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func genValue(n int) string {
	randomBytes := make([]byte, (n*3+1)/2)
	if _, err := io.ReadFull(rands.Reader, randomBytes); err != nil {
		log.Errorf("Error generating random value: %v", err)
		return ""
	}
	hexStr := hex.EncodeToString(randomBytes)
	return hexStr[0:n]
}

func NewRedisLock(store *redis.Client, key string) *RedisLock {
	return &RedisLock{
		store:   store,
		key:     key,
		id:      genValue(randomLen),
		seconds: 30,
	}
}

func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	val, err := lockScript.Run(ctx, rl.store, []string{rl.key}, rl.id, int64(seconds*millisPerSecond+tolerance)).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		log.Errorf("Error on acquiring lock for %s, %s", rl.key, err.Error())
		return false, err
	} else if val == nil {
		return false, nil
	}

	reply, ok := val.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	log.Errorf("Unknown reply when acquiring lock for %s: %v", rl.key, val)
	return false, nil
}

func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	val, err := delScript.Run(ctx, rl.store, []string{rl.key}, rl.id).Result()
	if err != nil {
		return false, err
	}

	reply, ok := val.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

func (rl *RedisLock) SetExpire(seconds uint32) {
	atomic.StoreUint32(&rl.seconds, seconds)
}
