package prompt

import (
	"os"

	"golang.org/x/term"
)

// see https://github.com/c-bata/go-prompt/issues/233

var termState *term.State

func saveState() {
	state, _ := term.GetState(int(os.Stdin.Fd()))
	termState = state
}

func restoreState() {
	if termState != nil {
		term.Restore(int(os.Stdin.Fd()), termState)
	}
}
