package cli

import (
	"github.com/spf13/cobra"
)

type CliInput struct {
	SourceDir      string
	DestinationDir string
	IsOverwrite    bool
	Includes       []string
}

func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
func (cli *CliInput) IsSourceDirSelected() bool {
	return cli.SourceDir != ""
}
func (cli *CliInput) IsDestinationDirSelected() bool {
	return cli.DestinationDir != ""
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	source, _ := cmd.Flags().GetString("source")
	destination, _ := cmd.Flags().GetString("dest")
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	filenames, _ := cmd.Flags().GetStringSlice("only")

	input := CliInput{
		SourceDir:      source,
		DestinationDir: destination,
		IsOverwrite:    overwrite,
		Includes:       filenames,
	}

	return input
}
