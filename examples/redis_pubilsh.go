package main

import (
	"github.com/redis/go-redis/v9"
	pusher "github.com/rosas99/monster/pkg/streams/connector/redis"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	redisOptions := redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	}

	pusherOptions := pusher.NewPusherOptions()
	rds := redis.NewClient(&redisOptions)
	ins := pusher.NewPusher(pusherOptions, rds)
	ins.Start()

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit

	ins.Stop()

}
