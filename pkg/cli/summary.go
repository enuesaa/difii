package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repo"
	"github.com/fatih/color"
)

type SummaryService struct{}

func (srv *SummaryService) Confirm(fsio repo.FsioInterface) bool {
	return fsio.Confirm("Would you like to show diffs summary?")
}

func (srv *SummaryService) Plan(fsio repo.FsioInterface, input CliInput) {
	fsio.Printf(color.HiWhiteString("----- Summary -----\n"))
	fsio.Printf(color.HiWhiteString("Any additions or deletions to [%s] are shown.\n", input.WorkDir))
	fsio.Printf("\n")
}

func (srv *SummaryService) Render(fsio repo.FsioInterface, input CliInput) {
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
	fsio.Printf("\n")
}
