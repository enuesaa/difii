package cli

import (
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
	ren.Printf("\n")
	ren.Printf("Inspecting diffs..\n")
	ren.Printf("\n")

	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	if input.IsFileSpecified() {
		sourcefiles = files.FilterFiles(sourcefiles, input.Includes)
	}

	for _, filename := range sourcefiles {
		source := files.ReadStream(input.CompareDir + "/" + filename)
		dest := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		ren.Printf(
			"%s has %s %s diffs\n",
			filename,
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
		)

		srv.renderHunks(ren, *diffs)
		ren.Printf("\n")
	}
}

func (srv *InspectService) renderHunks(ren RendererInterface, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			if item.Added() {
				ren.Printf("%s\n", color.GreenString("+ " + item.Text()))
			} else {
				ren.Printf("%s\n", color.RedString("- " + item.Text()))
			}
		}
	}
}
