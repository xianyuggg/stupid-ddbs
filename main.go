package main

import (
	"stupid-ddbs/internal/mongo"
)

func main() {

	if err := mongo.LoadAllDataToMongo(); err != nil {
		panic(err)
	}

	//redis.RedisConnectTest()
	//hdfs.HDFSConnectTest()

}
