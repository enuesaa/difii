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
	fsio.Printf(color.HiWhiteString("----- Inspect -----\n"))

	targetfiles := listTargetFiles(fsio, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		source := fsio.ReadStream(input.CompareDir + "/" + filename)
		dest := fsio.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		srv.renderHunks(fsio, filename, *diffs)
	}
}

func (srv *InspectService) renderHunks(fsio repo.FsioInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				fsio.Printf("%s\t%s\n", filename+":"+line, color.GreenString("+ "+item.Text()))
			} else {
				fsio.Printf("%s\t%s\n", filename+":"+line, color.RedString("- "+item.Text()))
			}
		}
	}
	if len(diffs.ListItems()) > 0 {
		fsio.Printf("\n")
	}
}
