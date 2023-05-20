package prompt

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
)

func promptOptions() []prompt.Option {
	options := make([]prompt.Option, 0)

	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind {
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionShowCompletionAtStart())
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionMaxSuggestion(15))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))
	options = append(options, prompt.OptionCompletionOnDown())
	options = append(options, prompt.OptionBreakLineCallback(func(d *prompt.Document) {
		fmt.Println("a")
	}))

	return options
}