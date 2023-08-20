package cli

import (
	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/fatih/color"
)

func Inspect(renderer RendererInterface, input CliInput) {
	renderer.Printf("Inspecting diffs..\n")
	renderer.Printf("\n")

	sourcefiles := files.ListFilesRecursively(input.CompareDir)

	if input.IsFileSpecified() {
		sourcefiles = files.FilterFiles(sourcefiles, input.Includes)
	}

	for _, filename := range sourcefiles {
		source := files.ReadStream(input.CompareDir + "/" + filename)
		dest := files.ReadStream(input.WorkDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		renderer.Printf(
			"%s has %s %s diffs\n",
			filename,
			color.RedString("-%d", diffs.CountRemove()),
			color.GreenString("+%d", diffs.CountAdd()),
		)

		renderWithHunks(renderer, *diffs)
	}
}

func renderWithHunks(renderer RendererInterface, diffs diff.Diffs) {
	for _, hunk := range diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			if item.Added() {
				renderer.Printf("%s\n", color.GreenString("+ " + item.Text()))
			} else {
				renderer.Printf("%s\n", color.RedString("- " + item.Text()))
			}
		}
	}
	renderer.Printf("\n")
}
