package cmd

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

var startSuggests = []prompt.Suggest{
	{
		Text:        "set",
		Description: "set display_details/sharding true/false",
	},
	{
		Text:        "ping",
		Description: "to check database connection",
	},
	{
		Text:        "query",
		Description: "to query a collection using a set of attribute",
	},
	{
		Text:        "show",
		Description: "to show hdfs/collections/shards",
	},
	{
		Text:        "load",
		Description: "to load collection",
	},
	{
		Text:        "drop",
		Description: "to drop collection",
	},
	{
		Text:        "exit",
		Description: "exit this program",
	},
}


var multiStepSuggests = map[string][]prompt.Suggest{
	"ping": {
	},
	"load": {
		{Text: "local"},
		{Text: "article"},
		{Text: "read"},
		{Text: "user"},
		{Text: "beread"},
		{Text: "popular"},
	},
	"drop": {
		{Text: "all"},
		{Text: "article"},
		{Text: "read"},
		{Text: "user"},
		{Text: "beread"},
		{Text: "popular"},
	},
	"show": {
		{Text: "hdfs"},
		{Text: "collections"},
		{Text: "shards"},
	},
	"set": {
		{Text: "display_details"},
		{Text: "sharding"},
	},
	"query": {
		{Text: "article"},
		{Text: "user"},
		{Text: "read"},
		{Text: "beread"},
	},
}

func solveMultiStepSuggests(
	in prompt.Document,
	now string,
	fields []string,
) []prompt.Suggest {
	suggests, found := multiStepSuggests[now]
	if !found {
		prompt.FilterHasPrefix([]prompt.Suggest{}, in.GetWordBeforeCursor(), true)
	}
	//
	for _, text := range fields {
		for _, suggest := range suggests {
			if text == suggest.Text {
				now = now + "." + text
				suggests, found = multiStepSuggests[now]
				prompt.FilterHasPrefix([]prompt.Suggest{}, in.GetWordBeforeCursor(), true)
			}
		}
	}
	return prompt.FilterHasPrefix(suggests, in.GetWordBeforeCursor(), true)
}

func completer(in prompt.Document) []prompt.Suggest {
	check := func(c rune) bool {
		return c == ' ' || c == '\n' || c == ';'
	}

	suffix := strings.ToLower(in.CurrentLine())

	fields := strings.FieldsFunc(suffix, check)

	if len(fields) > 0 && !check(rune(suffix[len(suffix)-1])) {
		fields = fields[0 : len(fields)-1]
	}

	if len(fields) == 0 {
		return prompt.FilterHasPrefix(startSuggests, in.GetWordBeforeCursor(), true)
	}

	return solveMultiStepSuggests(in, fields[0], fields)
}

