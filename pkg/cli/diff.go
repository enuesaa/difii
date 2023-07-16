package cli

import (
	"fmt"
	"io"
	"strings"

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

		renderHunks(diffs.ListHunks(), files.ReadStream(input.WorkDir + "/" + filename))
	}
}

func renderHunks(hunks []diff.Hunk, dest io.Reader) {
	// see https://forum.golangbridge.org/t/how-can-i-read-desired-line-from-the-file/4268/2
	raw, _ := io.ReadAll(dest)
	lines := strings.Split(string(raw), "\n")

	scanned := 0
	for _, hunk := range hunks {
		for i, item := range hunk.ListItems() {
			if i == 0 && len(lines) > item.Line() {
				fmt.Println(lines[item.Line() - 2])
			}
			if item.Added() {
				fmt.Println(color.GreenString("+ " + item.Text()))
			} else {
				fmt.Println(color.RedString("- " + item.Text()))
			}
			scanned = item.Line()
		}
		if len(lines) > scanned {
			fmt.Println(lines[scanned])
		}
	}
	fmt.Printf("\n")
}
