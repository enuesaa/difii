package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type ImportService struct{}

func (srv *ImportService) Confirm(fsio repo.FsioInterface) bool {
	return fsio.Confirm("Would you like to import diffs?")
}

func (srv *ImportService) Render(fsio repo.FsioInterface, input CliInput) {
	fsio.Printf(color.HiWhiteString("----- Import -----\n"))

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

func (srv *ImportService) renderHunks(fsio repo.FsioInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				fsio.Printf("%-10s %s\n", filename+":"+line, color.GreenString("+ "+item.Text()))
			} else {
				fsio.Printf("%-10s %s\n", filename+":"+line, color.RedString("- "+item.Text()))
			}
		}
		if fsio.Confirm("Would you like to import this hunk?") {
			srv.importHunk(fsio, filename, hunk)
		}
	}
}

func (srv *ImportService) importHunk(fsio repo.FsioInterface, filename string, hunk diff.Hunk) {
	fsio.Printf("%+v\n", hunk)
}