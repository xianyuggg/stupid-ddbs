package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	log "stupid-ddbs/logutil"
)

func Run() {
		fmt.Println("")
		fmt.Println("  ____    _____   _   _   ____    ___   ____            ____    ____    ____    ____ ")
		fmt.Println("/ ___|  |_   _| | | | | |  _ \\  |_ _| |  _ \\          |  _ \\  |  _ \\  | __ )  / ___|")
		fmt.Println(":___) |   | |   | |_| | |  __/   | |  | |_| | |_____| | |_| | | |_| | | |_) |  ___) |")
		fmt.Println("|____/    |_|    \\___/  |_|     |___| |____/          |____/  |____/  |____/  |____/")

	fmt.Println("")
		log.SetLevel(1)

	reader := prompt.New(executor, completer,
			prompt.OptionTitle("stupid-ddbs cmd"),
			prompt.OptionHistory(history),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
			prompt.OptionPrefix("stupid-ddbs >>> "),
			prompt.OptionLivePrefix(changeLivePrefix),
		)

	reader.Run()
}
