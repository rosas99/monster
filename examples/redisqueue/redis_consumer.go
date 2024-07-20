package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	stream "github.com/rosas99/monster/pkg/streams/connector/redis"
	"github.com/rosas99/monster/pkg/streams/flow"
	"k8s.io/apimachinery/pkg/util/wait"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addUTC = func(msg any) any {
	timestamp := time.Now().Format(time.DateTime)

	// Concatenate the UTC timestamp with msg.Value
	if value, ok := msg.(string); ok {
		msg = timestamp + " " + value
	}

	return msg
}

const redisChannel = "channel"

func main() {

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞

	stopCh := make(chan struct{})
	ctx := wait.ContextForChannel(stopCh)
	redisOptions := redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	}

	source, err := stream.NewRedisSource(ctx, &redisOptions, redisChannel)
	if err != nil {
		fmt.Println("redis source err:", err)
	}

	sink := stream.NewRedisSink(&redisOptions, redisChannel)

	filter := flow.NewMap(addUTC, 1)
	source.Via(filter).To(sink)

	<-quit
	stopCh <- struct{}{}

}
