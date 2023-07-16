package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func ShowDiffs(input CliInput) {
	fmt.Printf("Inspecting diffs..\n")
	fmt.Printf("  %s\n", input.WorkDir)
	fmt.Printf("  %s\n", input.CompareDir)
	fmt.Printf("\n")

	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	if input.IsFileSpecified() {
		sourcefiles = files.FilterFiles(sourcefiles, input.Includes)
	}

	for _, filename := range sourcefiles {
		source := files.ReadStream(input.CompareDir + "/" + filename)
		dest := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		fmt.Printf(
			"%s has %s %s diffs\n",
			filename,
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
		)

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
