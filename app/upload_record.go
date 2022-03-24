package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	cosv5 "github.com/tencentyun/cos-go-sdk-v5"

	"github.com/bubble-diff/bubblereplay/models"
)

func UploadRecord(ctx context.Context, record *models.Record) (cosKey string, err error) {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%d/%s", record.TaskID, UUID.String())

	opt := &cosv5.ObjectPutOptions{
		ObjectPutHeaderOptions: &cosv5.ObjectPutHeaderOptions{
			ContentType: "application/json",
		},
	}

	b, err := json.Marshal(record)
	if err != nil {
		return "", err
	}

	_, err = cos.Object.Put(ctx, key, bytes.NewReader(b), opt)
	if err != nil {
		return "", err
	}

	return UUID.String(), nil
}
