package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stupid-ddbs/logutil"
)

var DefaultDbName = "ProjectDB"

func mongoGetDatabase() (*mongo.Database, error){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:20053")

	// Connect to MongoDB
	if client, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Error(err)
		return nil, err
	} else {
		return client.Database(DefaultDbName), nil
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
		log.Error(err)
	}
	log.Info("connected")
}


