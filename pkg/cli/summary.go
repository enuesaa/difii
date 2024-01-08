package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

type SummaryService struct{}

func (srv *SummaryService) Plan(fsio repository.FsioInterface, input CliInput) {
	if !input.Interactive {
		return
	}
	fsio.Printf(color.HiWhiteString("----- Summary -----\n"))
	fsio.Printf(color.HiWhiteString("Any additions or deletions to [%s] are shown.\n", input.WorkDir))
	fsio.Printf("\n")
}

func (srv *SummaryService) Render(fsio repository.FsioInterface, input CliInput) {
	targetfiles := listTargetFiles(fsio, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		workDir := fsio.ReadStream(input.WorkDir + "/" + filename)
		compareDir := fsio.ReadStream(input.CompareDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		fsio.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
}
