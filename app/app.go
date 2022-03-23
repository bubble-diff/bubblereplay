package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bubble-diff/bubblereplay/config"
)

var (
	ctx = context.Background()
	cfg = config.Get()

	mongodb  *mongo.Client
	rdb      *redis.Client
	taskColl *mongo.Collection
)

func Init() (err error) {
	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	mongodb, err = mongo.Connect(ctx1, options.Client().ApplyURI(cfg.Mongo.Url))
	if err != nil {
		return err
	}

	rdb = redis.NewClient(&redis.Options{Addr: cfg.Redis.Addr, Password: cfg.Redis.Password})

	initColl()

	log.Println("init app ok.")
	return nil
}

func initColl() {
	taskColl = mongodb.Database(fmt.Sprintf("bubblediff_%s", cfg.Env)).Collection(cfg.Mongo.Collections.Task)
}
