package app

import (
	"context"
	"fmt"
	"time"

	"github.com/bubble-diff/bubblereplay/config"
)

// SetDeployed 设置diff任务对应的bubblecopy已部署
// example: bubblecopy_test_1 -> 127.0.0.1:8888
func SetDeployed(ctx context.Context, taskid int64, addr string, ttl time.Duration) (err error) {
	conf := config.Get()
	key := fmt.Sprintf("bubblecopy_%s_%d", conf.Env, taskid)
	val := addr
	return rdb.Set(ctx, key, val, ttl).Err()
}
