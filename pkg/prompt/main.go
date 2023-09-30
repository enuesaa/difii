package prompt

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


type MockPrompt struct {
	Out string
}
func NewMockPrompt() *MockPrompt {
	return &MockPrompt{}
}
func (prompt *MockPrompt) Printf(format string, a ...any) {
	prompt.Out += fmt.Sprintf(format, a...)
}