package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

func NewSummaryService(fsio repository.FsioInterface) SummaryService {
	return SummaryService{
		fsio: fsio,
	}
}

type SummaryService struct{
	fsio repository.FsioInterface
}

func (srv *SummaryService) Plan(input CliInput) {
	if !input.Interactive {
		return
	}
	srv.fsio.Printf(color.HiWhiteString("----- Summary -----\n"))
	srv.fsio.Printf(color.HiWhiteString("Any additions or deletions to [%s] are shown.\n", input.WorkDir))
	srv.fsio.Printf("\n")
}

func (srv *SummaryService) Render(input CliInput) {
	targetfiles := listTargetFiles(srv.fsio, input.WorkDir, input.CompareDir)
	if input.IsFileSpecified() {
		targetfiles = filterIncludeFiles(targetfiles, input.Includes)
	}

	for _, filename := range targetfiles {
		workDir := srv.fsio.ReadStream(input.WorkDir + "/" + filename)
		compareDir := srv.fsio.ReadStream(input.CompareDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		srv.fsio.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
}
