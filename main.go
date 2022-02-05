package main

import (
	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblereplay/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run("0.0.0.0:8080")
}
