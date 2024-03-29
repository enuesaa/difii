package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

type InspectService struct{}

func (srv *InspectService) Render(fsio repository.FsioInterface, input CliInput) {
	if input.Interactive {
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

func (srv *InspectService) renderHunks(fsio repository.FsioInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				fsio.Printf("%s:%-3s %s\n", filename, line, color.GreenString("+ "+item.Text()))
			} else {
				fsio.Printf("%s:%-3s %s\n", filename, line, color.RedString("- "+item.Text()))
			}
		}
	}
}
