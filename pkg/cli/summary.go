package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/repository"
	"github.com/fatih/color"
)

func NewSummaryService(repos repository.Repos) SummaryService {
	return SummaryService{
		fsio: repos.Fsio,
		log: repos.Log,
	}
}

type SummaryService struct{
	fsio repository.FsioInterface
	log repository.LogInterface
}

func (srv *SummaryService) Plan(input CliInput) {
	if !input.Interactive {
		return
	}
	srv.log.Printf(color.HiWhiteString("----- Summary -----\n"))
	srv.log.Printf(color.HiWhiteString("Any additions or deletions to [%s] are shown.\n", input.WorkDir))
	srv.log.Printf("\n")
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

		srv.log.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
}
