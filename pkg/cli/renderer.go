package cli

import (
	"fmt"
)

type RendererInterface interface {
	Printf(format string, a ...any)
}

type Renderer struct {}
func NewRenderer() *Renderer {
	return &Renderer{}
}
func (ren *Renderer) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

type MockRenderer struct {
	Out string
}
func NewMockRenderer() *MockRenderer {
	return &MockRenderer{}
}
func (ren *MockRenderer) Printf(format string, a ...any) {
	ren.Out += fmt.Sprintf(format, a...)
}
