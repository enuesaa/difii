package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func ShowDiffsSummary(input CliInput) {
	fmt.Printf("Diffs Summary\n")
	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	for _, filename := range sourcefiles {
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		workDir := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		fmt.Printf(
			"%12s %12s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
	fmt.Printf("\n")
}
