package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	log "stupid-ddbs/logutil"
)

func Run() {
		fmt.Println("")
		fmt.Println("   ▄████████     ███     ███    █▄     ▄███████▄  ▄█  ████████▄       ████████▄  ████████▄  ▀█████████▄     ▄████████ ")
		fmt.Println("  ███    ███ ▀█████████▄ ███    ███   ███    ███ ███  ███   ▀███      ███   ▀███ ███   ▀███   ███    ███   ███    ███ ")
		fmt.Println("  ███    █▀     ▀███▀▀██ ███    ███   ███    ███ ███▌ ███    ███      ███    ███ ███    ███   ███    ███   ███    █▀  ")
		fmt.Println("  ███            ███   ▀ ███    ███   ███    ███ ███▌ ███    ███      ███    ███ ███    ███  ▄███▄▄▄██▀    ███        ")
		fmt.Println("  ██████████     ███     ███    ███ ▀█████████▀  ███▌ ███    ███      ███    ███ ███    ███ ▀▀███▀▀▀██▄  ▀███████████ ")
		fmt.Println("         ███     ███     ███    ███   ███        ███  ███    ███      ███    ███ ███    ███   ███    ██▄          ███ ")
		fmt.Println("   ▄█    ███     ███     ███    ███   ███        ███  ███   ▄███      ███   ▄███ ███   ▄███   ███    ███    ▄█    ███ ")
		fmt.Println(" ▄████████▀     ▄████▀   ████████▀   ▄████▀      █▀   ████████▀       ████████▀  ████████▀  ▄█████████▀   ▄████████▀ ")
		fmt.Println("")

	log.SetLevel(0)

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
