package redislock

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	// 初始化 go-redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	// 创建 RedisLock 实例
	lock := NewRedisLock(rdb, "my_lock_key")
	lock.SetExpire(30) // 设置锁的超时时间为 30 秒

	// 获取锁
	if acquired, err := lock.AcquireCtx(context.Background()); acquired {
		defer func() { _, _ = lock.ReleaseCtx(context.Background()) }()
		// 安全地执行业务逻辑
		fmt.Println("Lock acquired, doing work...")
	} else {
		if err != nil {
			log.Printf("Failed to acquire lock: %v", err)
		} else {
			fmt.Println("Failed to acquire lock.")
		}
	}
}
