package main

import (
	"stupid-ddbs/internal/mongo"
	"stupid-ddbs/internal/printer"
)

func main() {

	//if err := mongo.LoadAllData(); err != nil {
	//	panic(err)
	//}

	manager := mongo.GetManagerInstance()
	res := manager.QueryData("article", []mongo.Cond{
		{"id", mongo.OpCompGE, "a1000"},
		{"id", mongo.OpCompLE, "a1003"},
		//{"title", mongo.OpCompEQ, "title1002"},
	})
	printer.ResultPrinter("article", res)

	//redis.RedisConnectTest()
	//hdfs.HDFSConnectTest()

}
