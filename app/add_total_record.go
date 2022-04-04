package app

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func AddTotalRecord(ctx context.Context, taskid int64) (err error) {
	filter := bson.D{{"id", taskid}}
	update := bson.D{
		{
			Key:   "$inc",
			Value: bson.D{{Key: "total_record", Value: 1}},
		},
	}

	result, err := taskColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Printf("add total record ok, taskid: %d, result: %+v", taskid, result)
	return nil
}
