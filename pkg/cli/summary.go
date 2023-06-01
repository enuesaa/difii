package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/olekukonko/tablewriter"
)

func Summary(input CliInput) {
	fmt.Printf("\nSummary..\n")
	sourcefiles := files.ListFilesRecursively(input.SourceDir)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"filename", "diffs", "source", "dest"})

	for _, filename := range sourcefiles {
		sourcePath := input.SourceDir + "/" + filename
		destPath := input.DestDir + "/" + filename
		source := files.ReadStream(sourcePath)
		dest := files.ReadStream(destPath)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()
		table.Append([]string{filename, diffs.Summary(), sourcePath, destPath})
	}
	table.Render()
}
