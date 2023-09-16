package cli

import "github.com/fatih/color"

func Plan(ren RendererInterface, input CliInput) {
	ren.Printf(color.HiWhiteString("----- Compare -----\n"))
	ren.Printf("I'll show you any additions or deletions in [%s] when compared to [%s].\n", input.WorkDir, input.CompareDir)
	ren.Printf("\n")
}
