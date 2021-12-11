package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"stupid-ddbs/logutil"
)

func bulkLoadDataToMongo(db *mongo.Database, collectionName string, values []interface{}) error{
	collection := db.Collection(collectionName)
	_, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("bulk load ok:", collectionName)
	return nil
}

func LoadAllDataToMongo() error{
	db, err := mongoGetDatabase()
	if err != nil {
		return err
	}
	articles, _ := LoadArticleDataFromLocal("article")
	reads, _ := LoadArticleDataFromLocal("read")
	users, _ := LoadArticleDataFromLocal("user")

	_ = bulkLoadDataToMongo(db, "article", articles)
	_ = bulkLoadDataToMongo(db, "read", reads)
	_ = bulkLoadDataToMongo(db, "user", users)


	return mongoCloseDatabase(db)
}

