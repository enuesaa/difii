package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

func NewInspectService(fsio repository.FsioInterface) InspectService {
	return InspectService{
		fsio: fsio,
	}
}

type InspectService struct{
	fsio repository.FsioInterface
}

func (srv *InspectService) Render(input CliInput) {
	if input.Interactive {
		srv.fsio.Printf(color.HiWhiteString("----- Inspect -----\n"))
	}

	targetfiles := listTargetFiles(srv.fsio, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		comparedir := srv.fsio.ReadStream(input.CompareDir + "/" + filename)
		workdir := srv.fsio.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(comparedir, workdir)
		diffs := analyzer.Analyze()

		srv.renderHunks(filename, *diffs)
	}
}

func (srv *InspectService) renderHunks(filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				srv.fsio.Printf("%s:%-3s %s\n", filename, line, color.GreenString("+ "+item.Text()))
			} else {
				srv.fsio.Printf("%s:%-3s %s\n", filename, line, color.RedString("- "+item.Text()))
			}
		}
	}
}
