package cli

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func Diff(input CliInput) {
	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Printf("\n")
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
			prompt.Input("Do you overwrite ? [Y/n] ", func (in prompt.Document) []prompt.Suggest {
				return make([]prompt.Suggest, 0)
			})
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
