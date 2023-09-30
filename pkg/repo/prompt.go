package repo

import (
	"fmt"
)

type PromptInterface interface {
	Printf(format string, a ...any)
}


type Prompt struct {}
func NewPrompt() *Prompt {
	return &Prompt{}
}
func (prompt *Prompt) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}
