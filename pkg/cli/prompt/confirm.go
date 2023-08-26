package prompt

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func Confirm(message string) bool {
	saveState()
	answer := prompt.Input(message + " (y/N) ", confirmSuggestion, confirmPromptOptions()...)
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" || answer == "Y" {
		restoreState()
		return true
	}
	restoreState()
	return false
}

func confirmSuggestion(in prompt.Document) []prompt.Suggest {
	return make([]prompt.Suggest, 0)
}

func confirmPromptOptions() []prompt.Option {
	options := make([]prompt.Option, 0)

	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))

	return options
}
