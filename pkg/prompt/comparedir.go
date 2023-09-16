package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/difii/pkg/files"
)

func SelectCompareDir() string {
	// see https://github.com/c-bata/go-prompt/issues/8
	saveState()
	for {
		dir := prompt.Input("Compare dir (--compare): ", selectDir, comparedirPromptOptions()...)
		if files.IsDirOrFileExist(dir) {
			restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func comparedirPromptOptions() []prompt.Option {
	options := make([]prompt.Option, 0)

	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionShowCompletionAtStart())
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionMaxSuggestion(15))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))
	options = append(options, prompt.OptionCompletionOnDown())

	// TODO: 選んでenterを押すと入力途中なのに完了してしまうのをなんとかしたい

	return options
}

func isDirNamedLikeTextExist(text string) bool {
	if text == "." || strings.HasSuffix(text, "/") {
		return false
	}
	return files.IsDirOrFileExist(text)
}

func appendSuggest(suggests []prompt.Suggest, path string) []prompt.Suggest {
	suggests = append(suggests, prompt.Suggest{
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

func filterGit(suggests []prompt.Suggest) []prompt.Suggest {
	ret := make([]prompt.Suggest, 0)
	for _, suggest := range suggests {
		if !strings.HasSuffix(suggest.Text, ".git") {
			ret = append(ret, suggest)
		}
	}
	return ret
}

func selectDir(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	suggests = appendSuggest(suggests, "./")
	suggests = appendSuggest(suggests, "../")

	text := in.Text
	searchDir := getSearchDir(text)
	basePath := getBasePath(text)

	for _, dir := range files.ListDirs(searchDir) {
		suggests = appendSuggest(suggests, basePath+dir)
	}

	if isDirNamedLikeTextExist(text) {
		for _, dir := range files.ListDirs(text) {
			suggests = appendSuggest(suggests, text+"/"+dir)
		}
	}

	suggests = filterGit(suggests)

	return prompt.FilterHasPrefix(suggests, text, false)
}
