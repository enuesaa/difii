package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

type InspectService struct {}

func (srv *InspectService) Confirm() bool {
	return prompt.Confirm("Would you like to inspect diffs?")
}

func (srv *InspectService) Render(ren RendererInterface, input CliInput) {
	ren.Printf("-----------\n")
	ren.Printf("\n")
	ren.Printf("Inspect\n")
	ren.Printf("\n")

	targetfiles := files.ListFilesInDirs(input.WorkDir, input.CompareDir)

	if input.IsFileSpecified() {
		targetfiles = files.FilterFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		source := files.ReadStream(input.CompareDir + "/" + filename)
		dest := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		srv.renderHunks(ren, filename, *diffs)
		ren.Printf("\n")
	}
}

func (srv *InspectService) renderHunks(ren RendererInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				ren.Printf("%s\t%s\n", color.HiWhiteString(filename + ":" + line), color.GreenString("+ " + item.Text()))
			} else {
				ren.Printf("%s\t%s\n", color.HiWhiteString(filename + ":" + line), color.RedString("- " + item.Text()))
			}
		}
	}
}
