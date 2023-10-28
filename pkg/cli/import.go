package cli

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

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

		srv.renderHunks(fsio, filename, *diffs, input)
	}
}

func (srv *ImportService) renderHunks(fsio repo.FsioInterface, filename string, diffs diff.Diffs, input CliInput) {
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
			srv.importHunk(fsio, filename, hunk, input)
		}
	}
}

func (srv *ImportService) importHunk(fsio repo.FsioInterface, filename string, hunk diff.Hunk, input CliInput) {
	path := fmt.Sprintf("%s/%s", input.WorkDir, filename)

	// create file if not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, _ := os.Create(path)
		defer f.Close()
	}

	f, _ := os.Open(path)
	defer f.Close()
	contentbyte, _ := io.ReadAll(f)
	lines := strings.Split(string(contentbyte), "\n")

	for _, item := range hunk.ListItems() {
		if item.Added() {
			lines = slices.Insert(lines, item.Line(), item.Text())
		} else {
			lines = slices.Delete(lines, item.Line() - 1, item.Line())
		}
	}

	if len(lines) == 0 {
		os.Remove(path)
	} else {
		content := strings.Join(lines, "\n")
		f, _ := os.Create(path)
		defer f.Close()
		f.Write([]byte(content))
	}
}
