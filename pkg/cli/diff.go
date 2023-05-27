package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
)

func Diff(input CliInput) {
	fmt.Printf("source dir: %s \n", input.SourceDir)
	fmt.Printf("destination dir: %s \n", input.DestinationDir)

	sourcefiles := files.ListFilesRecursively(input.SourceDir)
	for _, filename := range sourcefiles {
		fmt.Printf("\n\n")
		fmt.Printf("diff start: %s\n", filename)
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestinationDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()
		fmt.Println(diffs.Render())
	}

	// search source files
	// for file in files
	//    read source file,
	//    read dest file,
	//    print diff
	// 

	// sourcePath := input.SourceDir + "/theme.ts"
	// destPath := input.DestinationDir + "/theme.ts"

	// source := files.ReadStream(sourcePath)
	// dest := files.ReadStream(destPath)

	// analyzer := diff.NewAnalyzer(source, dest)
	// diffs := analyzer.Analyze()
	// fmt.Println(diffs.Render())
}
