package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	goprompt "github.com/c-bata/go-prompt"
	"golang.org/x/term"
	"github.com/enuesaa/difii/pkg/files"
)

type PromptInterface interface {
	Printf(format string, a ...any)
	Confirm(message string) bool
	SelectCompareDir() string
}

type Prompt struct {
	termState *term.State
}
func NewPrompt() *Prompt {
	return &Prompt{}
}
func (prompt *Prompt) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

// see https://github.com/c-bata/go-prompt/issues/8
// see https://github.com/c-bata/go-prompt/issues/233
func (prompt *Prompt) saveState() {
	state, _ := term.GetState(int(os.Stdin.Fd()))
	prompt.termState = state
}

func (prompt *Prompt) restoreState() {
	if prompt.termState != nil {
		term.Restore(int(os.Stdin.Fd()), prompt.termState)
	}
	prompt.termState = nil
}

func (prompt *Prompt) Confirm(message string) bool {
	prompt.saveState()

	suggestion := func (in goprompt.Document) []goprompt.Suggest {
		return make([]goprompt.Suggest, 0)
	}
	options := make([]goprompt.Option, 0)
	options = append(options, goprompt.OptionAddKeyBind(goprompt.KeyBind{
		Key: goprompt.ControlC,
		Fn: func(*goprompt.Buffer) {
			prompt.restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, goprompt.OptionSuggestionBGColor(goprompt.Black))
	options = append(options, goprompt.OptionScrollbarThumbColor(goprompt.Black))
	options = append(options, goprompt.OptionSuggestionTextColor(goprompt.White))
	options = append(options, goprompt.OptionSelectedSuggestionBGColor(goprompt.Black))
	options = append(options, goprompt.OptionSelectedSuggestionTextColor(goprompt.Cyan))
	options = append(options, goprompt.OptionPrefixTextColor(goprompt.Brown))

	answer := goprompt.Input(message+" (y/N) ", suggestion, options...)
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" || answer == "Y" {
		prompt.restoreState()
		return true
	}
	prompt.restoreState()
	return false
}

func (prompt *Prompt) SelectCompareDir() string {
	prompt.saveState()

	options := make([]goprompt.Option, 0)
	// TODO: 選んでenterを押すと入力途中なのに完了してしまうのをなんとかしたい
	options = append(options, goprompt.OptionAddKeyBind(goprompt.KeyBind{
		Key: goprompt.ControlC,
		Fn: func(*goprompt.Buffer) {
			prompt.restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, goprompt.OptionShowCompletionAtStart())
	options = append(options, goprompt.OptionSuggestionBGColor(goprompt.Black))
	options = append(options, goprompt.OptionScrollbarThumbColor(goprompt.Black))
	options = append(options, goprompt.OptionSuggestionTextColor(goprompt.White))
	options = append(options, goprompt.OptionSelectedSuggestionBGColor(goprompt.Black))
	options = append(options, goprompt.OptionSelectedSuggestionTextColor(goprompt.Cyan))
	options = append(options, goprompt.OptionMaxSuggestion(15))
	options = append(options, goprompt.OptionPrefixTextColor(goprompt.Brown))
	options = append(options, goprompt.OptionCompletionOnDown())

	for {
		dir := goprompt.Input("Compare dir (--compare): ", prompt.selectDir, options...)
		if files.IsDirOrFileExist(dir) {
			prompt.restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

// TODO: move to files interface
func (prompt *Prompt) isDirNamedLikeTextExist(text string) bool {
	if text == "." || strings.HasSuffix(text, "/") {
		return false
	}
	return files.IsDirOrFileExist(text)
}

func (prompt *Prompt) appendSuggest(suggests []goprompt.Suggest, path string) []goprompt.Suggest {
	suggests = append(suggests, goprompt.Suggest{
		Text: path,
	})
	return suggests
}

// TODO: move to files interface
func (prompt *Prompt) getSearchDir(text string) string {
	if strings.HasSuffix(text, "/") {
		return text
	}
	return filepath.Dir(text)
}

// TODO: move to files interface
func (prompt *Prompt) getBasePath(text string) string {
	base := ""
	if strings.Contains(text, "/") {
		base = filepath.Dir(text) + "/"
	}
	return base
}

// TODO: move to files interface
func (prompt *Prompt) filterGit(suggests []goprompt.Suggest) []goprompt.Suggest {
	ret := make([]goprompt.Suggest, 0)
	for _, suggest := range suggests {
		if !strings.HasSuffix(suggest.Text, ".git") {
			ret = append(ret, suggest)
		}
	}
	return ret
}

// TODO: move to files interface around suggestion
func (prompt *Prompt) selectDir(in goprompt.Document) []goprompt.Suggest {
	suggests := make([]goprompt.Suggest, 0)
	suggests = prompt.appendSuggest(suggests, "./")
	suggests = prompt.appendSuggest(suggests, "../")

	text := in.Text
	searchDir := prompt.getSearchDir(text)
	basePath := prompt.getBasePath(text)

	for _, dir := range files.ListDirs(searchDir) {
		suggests = prompt.appendSuggest(suggests, basePath+dir)
	}

	if prompt.isDirNamedLikeTextExist(text) {
		for _, dir := range files.ListDirs(text) {
			suggests = prompt.appendSuggest(suggests, text+"/"+dir)
		}
	}

	suggests = prompt.filterGit(suggests)

	return goprompt.FilterHasPrefix(suggests, text, false)
}
