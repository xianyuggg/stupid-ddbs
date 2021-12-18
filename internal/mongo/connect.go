package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stupid-ddbs/logutil"
)

var DefaultDbName = "ProjectDB"

func mongoGetDatabase(dbName string) (*mongo.Database, error){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	if client, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Error(err)
		return nil, err
	} else {
		log.Info("client get database", DefaultDbName)
		if dbName == "" {
			return client.Database(DefaultDbName), nil
		} else {
			return client.Database(dbName), nil
		}
	}
}
func mongoCloseDatabase(db *mongo.Database) error{
	client := db.Client()
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func MongoConnectTest() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		log.Error(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		println(err.Error())
	} else {
		println("mongo connected")
	}
}


