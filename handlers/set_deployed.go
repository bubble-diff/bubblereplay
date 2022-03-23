package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblereplay/app"
)

type setDeployedHandler struct {
	TaskID int64 `json:"task_id"`
	// Addr 基准服务地址
	Addr string `json:"addr"`
	// ttl 部署凭证有效时长
	ttl time.Duration
}

var SetDeployedHandler = &setDeployedHandler{
	ttl: 10 * time.Second,
}

func (h *setDeployedHandler) SetDeployed(c *gin.Context) {
	var err error

	err = c.BindJSON(h)
	if err != nil {
		log.Printf("[SetDeployed] bind json failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	log.Printf("TaskID=%d Addr=%s", h.TaskID, h.Addr)

	err = app.SetDeployed(c, h.TaskID, h.Addr, h.ttl)
	if err != nil {
		log.Printf("[SetDeployed] set deployed failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": nil,
	})
}
