package cli

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
)

func DoPrompt() {
	dirs := make([]string, 0)
	filenames, _ := os.ReadDir("./")
	for _, fileinfo := range filenames {
		if fileinfo.IsDir() {
			dirs = append(dirs, fileinfo.Name() + "/")
		}
	}

    in := prompt.Input("Select Dir ", func (in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)
		for _, dirname := range dirs {
			suggests = append(suggests, prompt.Suggest{ Text: dirname })
		}
		return prompt.FilterHasPrefix(suggests, in.GetWordBeforeCursor(), true)
	})
    fmt.Println("Your input: " + in)
}