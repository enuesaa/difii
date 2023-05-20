package prompt

import (
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

	return options
}