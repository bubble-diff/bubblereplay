package app

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblereplay/models"
)

func GetTaskDetail(taskid int64) (task *models.Task, err error) {
	task = new(models.Task)
	filter := bson.D{{"id", taskid}}
	err = taskColl.FindOne(ctx, filter).Decode(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
