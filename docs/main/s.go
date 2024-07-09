package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/core"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		core.WriteResponse(c, nil, gin.H{
			"code": 0,
			"msg":  "ok",
		})

	})
	r.Run("127.0.0.1:9494") // 监听并在 0.0.0.0:8080 上启动服务
}
