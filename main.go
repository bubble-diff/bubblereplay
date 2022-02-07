package main

import (
	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblereplay/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	_record := r.Group("/record")
	_record.POST("/add", handlers.AddRecord)
	r.Run("localhost:6789")
}
