package prompt

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

func selectDir(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = append(suggests, prompt.Suggest {
		Text: "./",
	})
	suggests = append(suggests, prompt.Suggest {
		Text: "../",
	})

	text := in.Text
	if text == "" {
		text = "./"
	}

	var searchDir string
	if strings.HasSuffix(text, "/") {
		// /の一つ手前で suggest したい
		searchDir = text
	} else {
		searchDir = filepath.Dir(text)
	}

	files, err := os.ReadDir(searchDir)
	if err != nil {
		return suggests
	}
	for _, f := range files {
		if f.IsDir() {
			var suggestion string
			if strings.Contains(text, "/") {
				suggestion = filepath.Dir(text) + "/" + f.Name() + "/"
			} else {
				suggestion = f.Name() + "/"
			}

			suggests = append(suggests, prompt.Suggest {
				Text: suggestion,
			})
		}
	}

	return prompt.FilterHasPrefix(suggests, text, false)
}
