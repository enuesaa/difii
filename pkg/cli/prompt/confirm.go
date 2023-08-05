package prompt

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func ConfirmToApply() bool {
	saveState()
	for {
		answer := prompt.Input("Override diffs ? [y/n]", func(in prompt.Document) []prompt.Suggest {
			return make([]prompt.Suggest, 0)
		})

		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}

		fmt.Printf("Error: Unknown answer found.")
	}
}
