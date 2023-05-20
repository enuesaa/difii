package prompt

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

func selectDir(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	text := in.Text

	searchDir := text
	if !strings.HasSuffix(text, "/") {
		searchDir = filepath.Dir(text)
	}

	files, err := os.ReadDir(searchDir)
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
