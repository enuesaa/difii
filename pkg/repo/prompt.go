package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	goprompt "github.com/c-bata/go-prompt"
	"golang.org/x/term"
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

	suggestion := func(in goprompt.Document) []goprompt.Suggest {
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
		dir := goprompt.Input("Compare dir (--compare): ", prompt.suggestDirs, options...)
		if prompt.isDirOrFileExist(dir) {
			prompt.restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func (prompt *Prompt) suggestDirs(in goprompt.Document) []goprompt.Suggest {
	suggests := make([]goprompt.Suggest, 0)
	suggests = append(suggests, goprompt.Suggest{Text: "./"})
	suggests = append(suggests, goprompt.Suggest{Text: "../"})

	text := in.Text

	searchDir := text
	if !strings.HasSuffix(text, "/") {
		searchDir = filepath.Dir(text)
	}
	basePath := ""
	if strings.Contains(text, "/") {
		basePath = filepath.Dir(text) + "/"
	}

	for _, dir := range prompt.listDirs(searchDir) {
		suggests = append(suggests, goprompt.Suggest{Text: basePath + dir})
	}

	if text != "." && !strings.HasSuffix(text, "/") && prompt.isDirOrFileExist(text) {
		for _, dir := range prompt.listDirs(text) {
			suggests = append(suggests, goprompt.Suggest{Text: text + "/" + dir})
		}
	}

	// filter .git
	suggests = slices.DeleteFunc(suggests, func(suggest goprompt.Suggest) bool {
		return strings.HasSuffix(suggest.Text, ".git")
	})

	return goprompt.FilterHasPrefix(suggests, text, false)
}

func (prompt *Prompt) listDirs(path string) []string {
	dirs := make([]string, 0)
	files, err := os.ReadDir(path)
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

func (prompt *Prompt) isDirOrFileExist(path string) bool {
	// see https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
