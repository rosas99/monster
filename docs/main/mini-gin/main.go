package main

import (
	"fmt"
	"github.com/rosas99/monster/docs/main/mini-gin/gin"
	"github.com/rosas99/monster/docs/main/mini-gin/middleware"
	"log"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	r := gin.New()
	r.Use(middleware.Logger()) // global middleware

	r.AddRoute("GET", "/", func(c *gin.Context) {
		responseData := ResponseData{Message: "Hello Gin"}
		c.JSON(http.StatusOK, responseData)
	})

	fmt.Println("Server is running on port 9999")
	log.Fatal(r.Run(":9999"))
}
