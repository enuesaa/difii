package cli

import (
	"errors"

	"github.com/enuesaa/difii/pkg/files"
	"github.com/spf13/cobra"
)

type CliInput struct {
	SourceDir      string
	DestinationDir string
	IsOverwrite    bool
	IsSummary      bool
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
func (cli *CliInput) IsSummaryOnly() bool {
	return cli.IsSummary
}
func (cli *CliInput) Validate() error {
	if !cli.IsSourceDirSelected() {
		return errors.New("required option --source is missing")
	}
	if !cli.IsDestinationDirSelected() {
		return errors.New("required option --dest is missing")
	}
	if !files.IsDirExist(cli.SourceDir) {
		return errors.New("invalid file path specified in --source")
	}
	if !files.IsDirExist(cli.DestinationDir) {
		return errors.New("invalid file path specified in --dest")
	}
	return nil
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	source, _ := cmd.Flags().GetString("source")
	destination, _ := cmd.Flags().GetString("dest")
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	summary, _ := cmd.Flags().GetBool("summary")
	filenames, _ := cmd.Flags().GetStringSlice("only")

	input := CliInput{
		SourceDir:      source,
		DestinationDir: destination,
		IsOverwrite:    overwrite,
		IsSummary:      summary,
		Includes:       filenames,
	}

	return input
}
