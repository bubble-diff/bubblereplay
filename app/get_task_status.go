package app

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTaskStatus(ctx context.Context, taskid int64) (isRunning bool, err error) {
	var task bson.M
	filter := bson.D{{"id", taskid}}
	err = taskColl.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return false, err
	}
	isRunning, ok := task["is_running"].(bool)
	if !ok {
		return false, errors.New("cannot convert to bool")
	}
	return isRunning, nil
}
