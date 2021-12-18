package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stupid-ddbs/logutil"
)

func CreateIndex(collName string, keyName string, unique bool) error{
	db, _:= mongoGetDatabase(DefaultDbName)
	coll := db.Collection(collName)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{
			keyName: 1,
		},
		Options: options.Index().SetCollation(&CollationConfig).SetUnique(unique),
	}
	if _, err := coll.Indexes().CreateOne(context.TODO(), indexModel, options.CreateIndexes()); err != nil {
		return err
	} else {
		return nil
	}
}

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

func ShardingSetup() {

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
	if err := db.CreateCollection(context.TODO(), "beread", nil); err != nil {
		log.Warning(err)
	}
	if err := db.CreateCollection(context.TODO(), "popular", nil); err != nil {
		log.Warning(err)
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "article")},
		{"key", bson.M{"category": 1}}, // 1 for range, "hashed" for hash
		//{"unique", true},
		//{"numInitialChunks", 32},  numInitialChunks is not supported when the shard key is not hashed.
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "read")},
		{"key", bson.M{"uid": 1}},
		//{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil {
		log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}


	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "user")},
		{"key", bson.M{"uid": 1}},
		//{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil{
		log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "beread")},
		{"key", bson.M{"aid": 1}},
		//{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	cmd = bson.D{
		{"shardCollection", fmt.Sprintf("%v.%v", DefaultDbName, "popular")},
		{"key", bson.M{"aid": 1}},
		//{"unique", true},
		//{"numInitialChunks", 32},
		{"collation",bson.M{"locale": "simple"}},
	}

	if err := db.RunCommand(context.TODO(), cmd).Err(); err != nil{
		log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}


	println("set up sharding finished")
	_ = mongoCloseDatabase(db)
}
