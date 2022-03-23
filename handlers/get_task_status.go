package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblereplay/app"
)

type getTaskStatusHandler struct{}

var GetTaskStatusHandler = &getTaskStatusHandler{}

// GetTaskStatus 获取diff任务是否运行
func (h *getTaskStatusHandler) GetTaskStatus(c *gin.Context) {
	taskid, err := strconv.ParseInt(c.Param("TaskID"), 10, 64)
	if err != nil {
		log.Printf("[GetTaskStatus] parse int failed, %s", err)
		c.JSON(200, gin.H{
			"err":        err.Error(),
			"is_running": false,
		})
		return
	}

	isRunning, err := app.GetTaskStatus(c, taskid)
	if err != nil {
		log.Printf("[GetTaskStatus] get task %d status failed, %s", taskid, err)
		c.JSON(200, gin.H{
			"err":        err.Error(),
			"is_running": false,
		})
		return
	}

	log.Printf("[GetTaskStatus] task=%d is_running=%v", taskid, isRunning)
	c.JSON(200, gin.H{
		"err":        "",
		"is_running": isRunning,
	})
}
