package cmd

import (
	"os"
	"strings"
	"stupid-ddbs/internal/mongo"
)

var livePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

var query string = ""

var history []string = make([]string, 0)

func executor(in string) {
	sql := strings.Trim(query+in, " \n;")
	if strings.HasSuffix(in, ";") || sql == "" {
		query += in
		livePrefixState.IsEnable = false
		livePrefixState.LivePrefix = in
		solve(sql)
		history = append(history, query)
		query = ""
		return
	}
	query += in + " "
	livePrefixState.LivePrefix = "... "
	livePrefixState.IsEnable = true
}

func changeLivePrefix() (string, bool) {
	return livePrefixState.LivePrefix, livePrefixState.IsEnable
}

// TODO: set global parameters
var displayDetails bool

func solve(query string) {
	if len(query) == 0 {
		return
	}
	commands := strings.Split(query, " ")
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "set":
		if len(commands) % 3 != 0 {
			println("set <attribute> value (display_details)")
			return
		}
		switch commands[1] {
		case "display_details":
			if commands[2] != "true" && commands[2] != "false" {
				println("value not valid (true/false)")
			}
			if commands[2] == "true" {
				displayDetails = true
			} else {
				displayDetails = false
			}
		default:
			println("attribute do not exist")
		}
	case "query":
		if (len(commands) - 2) % 3 != 0 {
			println("query <collection> {<attr> <lt> <value>}")
			return
		}
		collection := commands[1]
		conds := make([]mongo.Cond, 0)
		for i := 2; i < len(commands) - 2; i += 3 {
			attr := commands[i]
			op := mongo.OpCompGT
			switch commands[i+1] {
			case "eq":
				op = mongo.OpCompEQ
			case "lt":
				op = mongo.OpCompLT
			case "le":
				op = mongo.OpCompLE
			case "gt":
				op = mongo.OpCompGT
			case "ge":
				op = mongo.OpCompGE
			case "ne":
				op = mongo.OpCompNE
			default:
				println("unsupported op type")
				return
			}
			conds = append(conds, mongo.Cond{
				Field: attr,
				Op:    op,
				Val:   commands[i+2],
			})
		}
		mongo.QueryData(collection, conds, displayDetails)
	default:
		println("unknown command")
	}
}
