package main

import (
	"stupid-ddbs/internal/hdfs"
	"stupid-ddbs/internal/redis"
)

func main() {

	//mongo.LoadArticleDataFromLocal()
	redis.RedisConnectTest()
	hdfs.HDFSConnectTest()

}
