package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bubble-diff/bubblereplay/models"
)

func AppendRecordMeta(ctx context.Context, taskid int64, cosKey string, path string, rate float64) (err error) {
	key := fmt.Sprintf("task%d_records_meta", taskid)
	meta := models.RecordMeta{
		CosKey:   cosKey,
		Path:     path,
		DiffRate: fmt.Sprintf("%.2f%%", rate),
	}
	b, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	return rdb.LPush(ctx, key, b).Err()
}
