package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		print("收到")
		time.Sleep(5 * time.Second)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
