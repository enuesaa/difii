package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

func NewInspectService(repos repository.Repos) InspectService {
	return InspectService{
		fsio: repos.Fsio,
		log: repos.Log,
	}
}

type InspectService struct{
	fsio repository.FsioInterface
	log repository.LogInterface
}

func (srv *InspectService) Render(input CliInput) {
	if input.Interactive {
		srv.log.Printf(color.HiWhiteString("----- Inspect -----\n"))
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
				srv.log.Printf("%s:%-3s %s\n", filename, line, color.GreenString("+ "+item.Text()))
			} else {
				srv.log.Printf("%s:%-3s %s\n", filename, line, color.RedString("- "+item.Text()))
			}
		}
	}
}
