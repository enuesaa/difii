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
	ren.Printf("-----------\n")
	ren.Printf("\n")
	ren.Printf("Summary\n")
	ren.Printf("\n")

	targetfiles := files.ListFilesInDirs(input.BaseDir, input.CompareDir)

	if input.IsFileSpecified() {
		targetfiles = files.FilterFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		baseDir := files.ReadStream(input.BaseDir + "/" + filename)
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, baseDir)
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
