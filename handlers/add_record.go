package handlers

import "github.com/gin-gonic/gin"

func AddRecord(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add_record",
	})
}