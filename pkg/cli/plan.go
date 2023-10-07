package cli

import (
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

func Plan(fsio repo.FsioInterface, input CliInput) {
	fsio.Printf(color.HiWhiteString("----- Compare -----\n"))
	fsio.Printf("I'll show you any additions or deletions in [%s] when compared to [%s].\n", input.WorkDir, input.CompareDir)
	fsio.Printf("\n")
}
