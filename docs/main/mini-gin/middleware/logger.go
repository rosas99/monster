package middleware

import (
	"github.com/rosas99/monster/docs/main/mini-gin/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))

		//flag:=true
		//if flag {
		//	c.Fail(500, "error")
		//}
	}
}
