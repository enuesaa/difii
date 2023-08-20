package prompt

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func Confirm(message string) bool {
	saveState()
	answer := prompt.Input(message + " (y/N) ", func(in prompt.Document) []prompt.Suggest {
		return make([]prompt.Suggest, 0)
	})

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" || answer == "Y" {
		return true
	}
	return false
}
