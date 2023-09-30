package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type InspectService struct{}

func (srv *InspectService) Confirm(prompt repo.PromptInterface) bool {
	return prompt.Confirm("Would you like to inspect diffs?")
}

func (srv *InspectService) Render(prompt repo.PromptInterface, input CliInput) {
	prompt.Printf(color.HiWhiteString("----- Inspect -----\n"))

	targetfiles := files.ListFilesInDirs(input.WorkDir, input.CompareDir)

	if input.IsFileSpecified() {
		targetfiles = files.FilterFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		source := files.ReadStream(input.CompareDir + "/" + filename)
		dest := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		srv.renderHunks(prompt, filename, *diffs)
	}
}

func (srv *InspectService) renderHunks(prompt repo.PromptInterface, filename string, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			line := fmt.Sprint(item.Line())
			if item.Added() {
				prompt.Printf("%s\t%s\n", filename+":"+line, color.GreenString("+ "+item.Text()))
			} else {
				prompt.Printf("%s\t%s\n", filename+":"+line, color.RedString("- "+item.Text()))
			}
		}
	}
	if len(diffs.ListItems()) > 0 {
		prompt.Printf("\n")
	}
}
