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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit
	stopCh <- struct{}{}

}
