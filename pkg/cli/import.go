package cli

import (
	"fmt"
	"path/filepath"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type ImportService struct{}

func (srv *ImportService) Confirm(fsio repo.FsioInterface) bool {
	return fsio.Confirm("Would you like to import diffs? Warning: NOT STABLE")
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

		for _, hunk := range diffs.ListHunks() {
			srv.renderHunk(fsio, filename, hunk)
		}
		workfilePath := filepath.Join(input.WorkDir, filename)
		fsio.Printf("[%s] has diffs. \n", workfilePath)
		if fsio.Confirm("Would you like to overwrite this file?") {
			srv.importFile(fsio, filename, input)
		}
		fsio.Printf("\n")
	}
}

func (srv *ImportService) renderHunk(fsio repo.FsioInterface, filename string, hunk diff.Hunk) {
	for _, item := range hunk.ListItems() {
		line := fmt.Sprint(item.Line())
		if item.Added() {
			fsio.Printf("%-10s %s\n", filename+":"+line, color.GreenString("+ "+item.Text()))
		} else {
			fsio.Printf("%-10s %s\n", filename+":"+line, color.RedString("- "+item.Text()))
		}
	}
}

func (srv *ImportService) importFile(fsio repo.FsioInterface, filename string, input CliInput) error {
	workfilePath := filepath.Join(input.WorkDir, filename)
	comparefilePath := filepath.Join(input.CompareDir, filename)

	if err := fsio.RemoveFile(workfilePath); err != nil {
		return err
	}
	if err := fsio.CopyFile(comparefilePath, workfilePath); err != nil {
		return err
	}
	return nil
}
