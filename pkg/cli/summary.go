package cli

import (
	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

type SummaryService struct {}

func (srv *SummaryService) Confirm() bool {
	return prompt.Confirm("Would you like to show diffs summary?")
}

func (srv *SummaryService) Render(ren RendererInterface, input CliInput) {
	ren.Printf("Diffs Summary\n")
	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	if input.IsFileSpecified() {
		sourcefiles = files.FilterFiles(sourcefiles, input.Includes)
	}

	for _, filename := range sourcefiles {
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		workDir := files.ReadStream(input.WorkDir + "/" + filename)
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
