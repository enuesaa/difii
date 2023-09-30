package repo

import (
	"fmt"
)

type PromptMock struct {
	Out string
}

func NewPromptMock() *PromptMock {
	return &PromptMock{}
}
func (prompt *PromptMock) Printf(format string, a ...any) {
	prompt.Out += fmt.Sprintf(format, a...)
}
func (prompt *PromptMock) Confirm(message string) bool {
	return true
}
func (Prompt *PromptMock) SelectCompareDir() string {
	return "./"
}
