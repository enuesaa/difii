package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func ShowDiffs(input CliInput) {
	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Printf("%s\n", filename)
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		for _, hunk := range diffs.ListHunks() {
			for _, item := range hunk.ListItems() {
				renderDiffLine(item)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func renderDiffLine(item diff.Diffline) {
	if item.Added() {
		fmt.Println(color.GreenString("+ " + item.Text()))
	} else {
		fmt.Println(color.RedString("- " + item.Text()))
	}
}

// overwrite
