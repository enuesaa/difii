package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
)

func Diff(input CliInput) {
	fmt.Printf("source dir: %s \n", input.SourceDir)
	fmt.Printf("destination dir: %s \n", input.DestinationDir)
	fmt.Println("")

	sourcePath := input.SourceDir + "/theme.ts"
	destPath := input.DestinationDir + "/theme.ts"

	source := files.ReadStream(sourcePath)
	dest := files.ReadStream(destPath)

	analyzer := diff.NewAnalyzer(source, dest)
	diffs := analyzer.Analyze()
	fmt.Println(diffs.Render())
}
