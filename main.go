package main

import "stupid-ddbs/internal/mongo"

func main() {

//	fmt.Println("_____   _____   _   _   _____   _   _____        _____   _____   _____   _____")
//	fmt.Println("/  ___/ |_   _| | | | | |  _  \\ | | |  _  \\      |  _  \\ |  _  \\ |  _  \\ /  ___/")
//	fmt.Println("| |___    | |   | | | | | |_| | | | | | | |      | | | | | | | | | |_| | | |___")
//	fmt.Println("\\___  \\   | |   | | | | |  ___/ | | | | | |      | | | | | | | | |  _  { \\___  \\")
//	fmt.Println(" ___| |   | |   | |_| | | |     | | | |_| |      | |_| | | |_| | | |_| |  ___| |")
//	fmt.Println("/_____/   |_|   \\_____/ |_|     |_| |_____/      |_____/ |_____/ |_____/ /_____/")
//	fmt.Println("")
//
//
//
//reader := prompt.New(cmd.Executor, cmd.Completer,
//		prompt.OptionTitle("sql-prompt"),
//		prompt.OptionHistory(cmd.History),
//		prompt.OptionPrefixTextColor(prompt.Yellow),
//		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
//		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
//		prompt.OptionSuggestionBGColor(prompt.DarkGray),
//		prompt.OptionPrefix("stupid-ddbs >>> "),
//		prompt.OptionLivePrefix(cmd.ChangeLivePrefix),
//	)
//
//	reader.Run()

	manager := mongo.GetManagerInstance()

	//if err := manager.LoadAllData(); err != nil {
	//	panic(err)
	//}

	res := manager.QueryData("beread", []mongo.Cond{
		{"aid", mongo.OpCompGE, "1000"},
		{"aid", mongo.OpCompLE, "1100"},
		//{"title", mongo.OpCompEQ, "title1002"},
	})
	mongo.ResultPrinter("article", res)
	_ = manager.ComputeBeRead(false)
	_ = manager.QueryPopularRank()
	//mongo.ShardingSetup()


	//redis.RedisConnectTest()
	//hdfs.HDFSConnectTest()

}
