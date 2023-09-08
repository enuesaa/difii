package cli

import (
	// "fmt"
)

func Plan(ren RendererInterface, input CliInput) {
	ren.Printf("Comparing..\n")
	ren.Printf("\n")
	ren.Printf("- %s\n", input.WorkDir)
	ren.Printf("- %s\n", input.CompareDir)
	ren.Printf("\n")

	ren.Printf("I'll show you any additions or deletions in [%s] when compared to [%s].", input.WorkDir, input.CompareDir)
	ren.Printf("\n")
	ren.Printf("\n")
}
