package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli/prompt"
)

func Apply(input CliInput) {
	if input.Interactive && !prompt.ConfirmToApply() {
		return
	}
	fmt.Printf("applying...")
}
