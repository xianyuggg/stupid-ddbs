package main

import (
	"stupid-ddbs/internal/mongo"
)

func main() {

	//if err := mongo.LoadAllData(); err != nil {
	//	panic(err)
	//}

	mongo.ShardingSetup()

	//redis.RedisConnectTest()
	//hdfs.HDFSConnectTest()

}
