package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func ShowDiffsSummary(renderer RendererInterface, input CliInput) {
	renderer.Printf("Diffs Summary\n")
	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	if input.IsFileSpecified() {
		sourcefiles = files.FilterFiles(sourcefiles, input.Includes)
	}

	for _, filename := range sourcefiles {
		compareDir := files.ReadStream(input.CompareDir + "/" + filename)
		workDir := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(compareDir, workDir)
		diffs := analyzer.Analyze()

		renderer.Printf(
			"%s %s diffs in %s \n",
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
			filename,
		)
	}
	renderer.Printf("\n")
}
