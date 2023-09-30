package repo

import (
	"fmt"
)

type MockPrompt struct {
	Out string
}
func NewMockPrompt() *MockPrompt {
	return &MockPrompt{}
}
func (prompt *MockPrompt) Printf(format string, a ...any) {
	prompt.Out += fmt.Sprintf(format, a...)
}
func (prompt *MockPrompt) Confirm(message string) bool {
	return true
}
