package cli

import (
	"github.com/enuesaa/difii/pkg/repo"
)

func Plan(fsio repo.FsioInterface, input CliInput) {
	fsio.Printf("I'll show you any additions or deletions of [%s] when compared to [%s].\n", input.WorkDir, input.CompareDir)
	fsio.Printf("\n")
}
