package main

import (
	"stupid-ddbs/cmd"
	"stupid-ddbs/internal/mongo"
)

func main() {

	// time test
	//begin := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	//end := time.Date(2004, 12, 31, 23, 59, 0, 0, time.Local).Unix()
	//println(begin)
	//println(end)
	//println(1. * (begin - end)/1000000)
	//
	//println(time.Unix(0, 0).Format(time.RFC3339Nano))
	//os.Exit(0)

	//
	cmd.Run()
	//moniter.ShowHdfsPathStatus("/")
	//moniter.PrintShards()
	//moniter.PrintAllCollectionsStats()
	//mongo.PrintDbStats()

	manager := mongo.GetManagerInstance()

	////if err := manager.LoadAllData(); err != nil {
	////	panic(err)
	////}
	//
	_, _ = manager.QueryData("article", []mongo.Cond{
		{"aid", mongo.OpCompGE, "1000"},
		{"aid", mongo.OpCompLE, "1005"},
		{"title", mongo.OpCompEQ, "title1002"},
	})
	//mongo.ResultPrinter("article", res, false)
	//_ = manager.ComputeBeRead(false)
	//_ = manager.ComputePopular()
	//mongo.ShardingSetup()


	//redis.RedisConnectTest()
	//hdfsManager := hdfs.GetManagerInstance()
	//hdfsManager.PathInit()
	//hdfsManager.LoadDataIntoHDFS()
}
