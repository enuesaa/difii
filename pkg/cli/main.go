package cli

import (
	// "fmt"

	"github.com/spf13/cobra"
)

type CliInput struct {
	SourceDir string
	DestinationDir string
	IsOverwrite bool
	Includes []string
}
func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
// なんか期待通りじゃない
func (cli *CliInput) IsSourceDirSelected() bool {
	return &cli.SourceDir != nil
}
func (cli *CliInput) IsDestinationDirSelected() bool {
	return &cli.DestinationDir != nil
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	input := CliInput{}
	if len(args) == 1 {
		input.SourceDir = args[0]
	}
	if len(args) == 2 {
		input.SourceDir = args[0]
		input.DestinationDir = args[1]
	}
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	input.IsOverwrite = overwrite

	filenames, _ := cmd.Flags().GetStringSlice("only")
	input.Includes = filenames

	return input
}
