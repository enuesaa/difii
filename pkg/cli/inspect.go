package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type InspectService struct{}

func (srv *InspectService) Confirm(fsio repo.FsioInterface) bool {
	return fsio.Confirm("Would you like to inspect diffs?")
}

func (srv *InspectService) Render(fsio repo.FsioInterface, input CliInput) {
	if input.Verbose {
		fsio.Printf(color.HiWhiteString("----- Inspect -----\n"))
	}

	targetfiles := listTargetFiles(fsio, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		comparedir := fsio.ReadStream(input.CompareDir + "/" + filename)
		workdir := fsio.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(comparedir, workdir)
		diffs := analyzer.Analyze()

		srv.renderHunks(fsio, filename, *diffs)
	}
}

func (srv *InspectService) renderHunks(fsio repo.FsioInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				fsio.Printf("%-10s %s\n", filename+":"+line, color.GreenString("+ "+item.Text()))
			} else {
				fsio.Printf("%-10s %s\n", filename+":"+line, color.RedString("- "+item.Text()))
			}
		}
	}
	if len(diffs.ListItems()) > 0 {
		fsio.Printf("\n")
	}
}
