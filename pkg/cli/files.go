package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
)

func DiffFiles(sourceDir string, destinationDir string) {
	fmt.Printf("source dir: %s \n", sourceDir)
	fmt.Printf("destination dir: %s \n", destinationDir)
	fmt.Println("")

	sourcefiles := files.ListFilesRecursively(sourceDir)
	files.ReadStreamWithDiff(sourceDir, destinationDir, sourcefiles[0])
}

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
