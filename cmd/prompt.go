package cmd

import (
	"os"
	"strings"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

var query string = ""

var History []string = make([]string, 0)

func Executor(in string) {
	sql := strings.Trim(query+in, " \n;")
	if strings.HasSuffix(in, ";") || sql == "" {
		query += in
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in

		solve(sql)

		History = append(History, query)
		query = ""
		return
	}
	query += in + " "
	LivePrefixState.LivePrefix = "... "
	LivePrefixState.IsEnable = true
}

func ChangeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func solve(sqls string) {
	if sqls == "exit" {
		os.Exit(0)
	}
}
