package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/enuesaa/difii/pkg/prompt"
	"github.com/fatih/color"
)

type SummaryService struct{}

func (srv *SummaryService) Confirm() bool {
	return prompt.Confirm("Would you like to show diffs summary?")
}

func (srv *SummaryService) Render(ren RendererInterface, input CliInput) {
	ren.Printf(color.HiWhiteString("----- Summary -----\n"))

	targetfiles := files.ListFilesInDirs(input.WorkDir, input.CompareDir)

	if input.IsFileSpecified() {
		targetfiles = files.FilterFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		workDir := files.ReadStream(input.WorkDir + "/" + filename)
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		ren.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
	ren.Printf("\n")
}
