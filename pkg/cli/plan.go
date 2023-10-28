package cli

import (
	"github.com/enuesaa/difii/pkg/repo"
)

func Plan(fsio repo.FsioInterface, input CliInput) {
	fsio.Printf("Any additions or deletions are shown below. [%s] is considered the truth.\n", input.CompareDir)
	fsio.Printf("\n")
}
