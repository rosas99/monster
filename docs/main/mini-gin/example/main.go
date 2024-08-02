package main

import (
	"fmt"
	"github.com/rosas99/monster/docs/main/mini-gin/gin"
	"log"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.AddRoute("GET", "/", func(c *gin.Context) {
		responseData := ResponseData{Message: "Hello Gin"}
		c.JSON(http.StatusOK, responseData)
	})

	fmt.Println("Server is running on port 9999")
	err := r.Run(":9999")

	if err != nil {
		log.Fatal(err)
	}
}
