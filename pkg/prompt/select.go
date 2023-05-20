package prompt

import (
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/difii/pkg/files"
)

func isDirNamedLikeTextExist(text string) bool {
	if text == "." || strings.HasSuffix(text, "/") {
		return false
	}
	return files.IsDirExist(text)
}

func appendSuggest(suggests []prompt.Suggest, path string) []prompt.Suggest {
	suggests = append(suggests, prompt.Suggest {
		Text: path,
	})
	return suggests
}

func getSearchDir(text string) string {
	if strings.HasSuffix(text, "/") {
		return text
	}
	return filepath.Dir(text)
}

func getBasePath(text string) string {
	base := ""
	if strings.Contains(text, "/") {
		base = filepath.Dir(text) + "/"
	}
	return base
}

func selectDir(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = appendSuggest(suggests, "./")
	suggests = appendSuggest(suggests, "../")

	text := in.Text
	searchDir := getSearchDir(text)
	basePath := getBasePath(text)

	for _, dir := range files.ListDirs(searchDir) {
		suggests = appendSuggest(suggests, basePath + dir)
	}

	if isDirNamedLikeTextExist(text) {
		for _, dir := range files.ListDirs(text) {
			suggests = appendSuggest(suggests, text + "/" + dir)
		}
	}

	return prompt.FilterHasPrefix(suggests, text, false)
}
