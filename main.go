package main

import "stupid-ddbs/cmd"

func main() {

	cmd.Run()
	//moniter.ShowHdfsPathStatus("/")
	//moniter.PrintShards()
	//moniter.PrintAllCollectionsStats()
	//mongo.PrintDbStats()

	//manager := mongo.GetManagerInstance()

	////if err := manager.LoadAllData(); err != nil {
	////	panic(err)
	////}
	//
	//res, _ := manager.QueryData("article", []mongo.Cond{
	//	{"aid", mongo.OpCompGE, "1000"},
	//	{"aid", mongo.OpCompLE, "1001"},
	//	//{"title", mongo.OpCompEQ, "title1002"},
	//})
	//mongo.ResultPrinter("article", res, false)
	//_ = manager.ComputeBeRead(false)
	//_ = manager.QueryPopularRank()
	//mongo.ShardingSetup()


	//redis.RedisConnectTest()
	//hdfsManager := hdfs.GetManagerInstance()
	//hdfsManager.PathInit()
	//hdfsManager.LoadDataIntoHDFS()
}
