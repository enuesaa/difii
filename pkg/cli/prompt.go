package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

func SelectSourceDir() string {
	in := prompt.Input("Select Source Dir ", func (in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)
		text := in.Text

		if strings.HasSuffix(text, "/") {
			files, err := os.ReadDir(text)
			if err != nil {
				return suggests
			}
			for _, f := range files {
				if f.IsDir() {
					suggests = append(suggests, prompt.Suggest {
						Text: filepath.Dir(text) + "/" + f.Name(),
					})
				}
			}
			return suggests
		} else {
			files, err := os.ReadDir(filepath.Dir(text))
			if err != nil {
				return suggests
			}
			for _, f := range files {
				if f.IsDir() {
					if strings.Contains(text, "/") {
						suggests = append(suggests, prompt.Suggest {
							Text: filepath.Dir(text) + "/" + f.Name() + "/",
						})
					} else {
						suggests = append(suggests, prompt.Suggest {
							Text: f.Name() + "/",
						})
					}
				}
			}
			return prompt.FilterHasPrefix(suggests, text, false)
		}
	},
	prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			os.Exit(0)
		},
	}),
	prompt.OptionShowCompletionAtStart(),
	)

	return in
}

func SelectDestinationDir() string {
	suggests := make([]prompt.Suggest, 0)

    in := prompt.Input("Select Destination Dir ", func (in prompt.Document) []prompt.Suggest {
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
	return in
}
