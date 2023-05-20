package prompt

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

func listDirs(dir string) []string {
	dirs := make([]string, 0)
	files, err := os.ReadDir(dir)
	if err != nil {
		return dirs
	}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}

func appendDefaultSuggests(suggests []prompt.Suggest) []prompt.Suggest {
	suggests = append(suggests, prompt.Suggest {
		Text: "./",
	})
	suggests = append(suggests, prompt.Suggest {
		Text: "../",
	})

	return suggests
}

func selectDir(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = appendDefaultSuggests(suggests)

	text := in.Text

	var searchDir string
	if strings.HasSuffix(text, "/") {
		searchDir = text
	} else {
		searchDir = filepath.Dir(text)
	}

	for _, dir := range listDirs(searchDir) {
		var suggestion string
		if strings.Contains(text, "/") {
			suggestion = filepath.Dir(text) + "/" + dir
		} else {
			suggestion = dir
		}
		suggests = append(suggests, prompt.Suggest {
			Text: suggestion,
		})
	}

	// /の一つ手前で suggest したい
	if text != "." && !strings.HasSuffix(text, "/") {
		// see https://gist.github.com/mattes/d13e273314c3b3ade33f
		if _, err := os.Stat(text); !os.IsNotExist(err) {
			for _, dir := range listDirs(text) {
				suggestion := text + "/" + dir
				suggests = append(suggests, prompt.Suggest {
					Text: suggestion,
				})
			}
		}
	}

	return prompt.FilterHasPrefix(suggests, text, false)
}
