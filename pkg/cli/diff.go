package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func Diff(input CliInput) {
	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Println("")
		fmt.Printf("%s\n", filename)
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		for _, hunk := range diffs.ListHunks() {
			fmt.Println("")
			for _, item := range hunk.ListItems() {
				renderDiffLine(item)
			}
		}
	}
}

func renderDiffLine(item diff.Diffline) {
	if item.Added() {
		fmt.Println(color.GreenString("+ " + item.Text()))
	} else {
		fmt.Println(color.RedString("- " + item.Text()))
	}
}

func DiffTable(input CliInput) {
	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Println("")
		fmt.Printf("%s\n", filename)
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"source", "dest"})

		for _, hunk := range diffs.ListHunks() {
			for _, item := range hunk.ListItems() {
				if item.Added() {
					table.Append([]string{
						color.GreenString("+ " + item.Text()),
						"",
					})
				} else {
					table.Append([]string{
						"",
						color.RedString("- " + item.Text()),
					})
				}
			}
		}
		table.Render()
	}
}
