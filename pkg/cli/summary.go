package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type SummaryService struct{}

func (srv *SummaryService) Confirm(prompt repo.PromptInterface) bool {
	return prompt.Confirm("Would you like to show diffs summary?")
}

func (srv *SummaryService) Render(prompt repo.PromptInterface, files repo.FilesInterface, input CliInput) {
	prompt.Printf(color.HiWhiteString("----- Summary -----\n"))

	targetfiles := listTargetFiles(files, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		workDir := files.ReadStream(input.WorkDir + "/" + filename)
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		prompt.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
	prompt.Printf("\n")
}
