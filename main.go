package main

import (
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblereplay/app"
	"github.com/bubble-diff/bubblereplay/config"
	"github.com/bubble-diff/bubblereplay/handlers"
)

func main() {
	conf := config.Get()
	err := app.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.POST("/record/add", handlers.AddRecord)
	r.GET("/task_status/:taskid", handlers.GetTaskStatusHandler.GetTaskStatus)
	r.POST("/deploy", handlers.SetDeployedHandler.SetDeployed)

	pprof.Register(r)

	err = r.Run(conf.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}
}
