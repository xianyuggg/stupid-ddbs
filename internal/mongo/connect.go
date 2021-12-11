package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stupid-ddbs/logutil"
)

var DefaultDbName = "ProjectDB"

func mongo_get_database() (*mongo.Database, error){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	if client, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		return nil, err
	} else {
		return client.Database(DefaultDbName), nil
	}
}
func mongo_close_database(db *mongo.Database) {
	client := db.Client();
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
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
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}


