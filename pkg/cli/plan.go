package cli

import (
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

func Plan(prompt repo.PromptInterface, input CliInput) {
	prompt.Printf(color.HiWhiteString("----- Compare -----\n"))
	prompt.Printf("I'll show you any additions or deletions in [%s] when compared to [%s].\n", input.WorkDir, input.CompareDir)
	prompt.Printf("\n")
}
