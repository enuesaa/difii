package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/olekukonko/tablewriter"
)

func Summary(input CliInput) {
	fmt.Printf("Source dir: %s \n", input.SourceDir)
	fmt.Printf("Destination dir: %s \n", input.DestinationDir)
	fmt.Printf("\n")
	fmt.Println("Summary")
	sourcefiles := files.ListFilesRecursively(input.SourceDir)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"diffs", "filename"})

	for _, filename := range sourcefiles {
		source := files.ReadStream(input.SourceDir + "/" + filename)
		dest := files.ReadStream(input.DestinationDir + "/" + filename)
		analyzer := diff.NewAnalyzer(source, dest)
		diffs := analyzer.Analyze()
		table.Append([]string{diffs.Summary(), filename})
	}
	table.Render()
}
