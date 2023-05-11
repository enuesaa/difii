package cli

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
)

func DoPrompt() {
	suggests := make([]prompt.Suggest, 0)

    in := prompt.Input("Select Dir ", func (in prompt.Document) []prompt.Suggest {
		var dir string
		if in.Text == "" {
			dir = "./"
		} else {
			dir = in.Text
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			return prompt.FilterHasPrefix(suggests, in.GetWordBeforeCursor(), true)
		}
		suggests = nil
		suggests := make([]prompt.Suggest, 0)
		suggests = append(suggests, prompt.Suggest{ Text: "../" + in.Text })

		for _, fileinfo := range files {
			if fileinfo.IsDir() {
				suggests = append(suggests, prompt.Suggest{ Text: dir + fileinfo.Name() + "/" })
			}
		}
		return prompt.FilterHasPrefix(suggests, in.GetWordBeforeCursor(), true)
	})
    fmt.Println("Your input: " + in)
}