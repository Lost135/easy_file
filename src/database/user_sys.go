package database

import (
	"context"
	"easy_file/src/config"
	"easy_file/src/structs"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbCollect struct {
	database   string
	collection string
}

func DbSql(database, collection string) *DbCollect {

	return &DbCollect{
		database,
		collection,
	}
}

func (dc *DbCollect) IfExist(user structs.User) (count int64, res string, err error) {
	coll := config.DbClient.Database(dc.database).Collection(dc.collection)
	var result bson.M
	err = coll.FindOne(
		context.TODO(),
		bson.D{{"name", user.Name}},
		options.FindOne().SetProjection(bson.D{{"_id", 0}, {"name", 1}}),
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, "无用户", errors.New("无用户")
		}
		config.CatchErr(err)
	}
	return 1, "用户已存在", errors.New("用户名已存在")
}

func (dc *DbCollect) InsertOne(user structs.User) (count int64, res string, err error) {
	coll := config.DbClient.Database(dc.database).Collection(dc.collection)
	var result bson.M
	err = coll.FindOne(
		context.TODO(),
		bson.D{{"name", user.Name}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, "无此用户", nil
		}
		config.CatchErr(err)
	}
	return 1, "用户已存在", errors.New("用户名已存在")
}
