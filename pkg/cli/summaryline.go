package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func Summaryline(input CliInput) {
	fmt.Printf("\n")
	fmt.Printf("Diffs Summary\n")
	sourcefiles := files.ListFilesRecursively(input.SourceDir)

	for _, filename := range sourcefiles {
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()
		fmt.Printf("%12s %12s diffs in %s \n", color.RedString("-%d", diffs.CountRemove()), color.GreenString("+%d", diffs.CountAdd()), filename)
	}
}
