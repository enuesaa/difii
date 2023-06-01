package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
)

func Diff(input CliInput) {
	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Printf("\n")
		fmt.Printf("%s\n", filename)
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestinationDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()

		renderer := diff.NewHunkedRenderer(*diffs)
		renderer.Render()
	}

	// ハンクにまとめて
	// overwrite prompt
}
