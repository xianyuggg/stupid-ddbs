package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"stupid-ddbs/logutil"
)

func ShardingSetup() {
	//db, err := mongoGetDatabase("")
	//defer mongoCloseDatabase(db)
	//
	//var coll *mongo.Collection
	//coll = db.Collection("article")
	//indexName, err := coll.Indexes().CreateOne(
	//	context.Background(),
	//	mongo.IndexModel{
	//		Keys: bson.M{
	//			"id": 1,
	//		},
	//		Options: options.Index().SetUnique(true),
	//	},
	//)

	//db.col.find({"likes": {$gt:50}, $or: [{"by": "菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty(
	db, err := mongoGetDatabase("admin")
	if err != nil {
		log.Error(err)
		return
	}

	var cmd bson.D
	cmd = bson.D{
		{
			"enableSharding", DefaultDbName,
		},
	}
	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err)
	}
	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "article")},
		{"key", bson.M{"_id": 1}}, // 1 for range, "hashed" for hash
		{"unique", true},
		//{"numInitialChunks", 32},  numInitialChunks is not supported when the shard key is not hashed.
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err)
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "read")},
		{"key", bson.M{"_id": 1}},
		{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err)
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "user")},
		{"key", bson.M{"_id": 1}},
		{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err)
	}
	log.Info("set up sharding finish")
	_ = mongoCloseDatabase(db)
}
