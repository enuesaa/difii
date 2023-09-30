package repo

import (
	"fmt"
	"os"
	"strings"

	goprompt "github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type PromptInterface interface {
	Printf(format string, a ...any)
	Confirm(message string) bool
}

type Prompt struct {
	// see https://github.com/c-bata/go-prompt/issues/233
	termState *term.State
}
func NewPrompt() *Prompt {
	return &Prompt{}
}
func (prompt *Prompt) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

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
	answer := goprompt.Input(message+" (y/N) ", prompt.confirmSuggestion, prompt.confirmPromptOptions()...)
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" || answer == "Y" {
		prompt.restoreState()
		return true
	}
	prompt.restoreState()
	return false
}

func (prompt *Prompt) confirmSuggestion(in goprompt.Document) []goprompt.Suggest {
	return make([]goprompt.Suggest, 0)
}

func (prompt *Prompt) confirmPromptOptions() []goprompt.Option {
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

	return options
}
