package cli

import (
	"errors"

	"github.com/enuesaa/difii/pkg/files"
	"github.com/spf13/cobra"
)

type CliInput struct {
	SourceDir   string
	DestDir     string
	IsOverwrite bool
	Includes    []string
}

func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
func (cli *CliInput) IsSourceDirSelected() bool {
	return cli.SourceDir != ""
}
func (cli *CliInput) IsDestDirSelected() bool {
	return cli.DestDir != ""
}
func (cli *CliInput) Validate() error {
	if !cli.IsSourceDirSelected() {
		return errors.New("required option --source is missing")
	}
	if !cli.IsDestDirSelected() {
		return errors.New("required option --dest is missing")
	}
	if !files.IsDirExist(cli.SourceDir) {
		return errors.New("invalid file path specified in --source")
	}
	if !files.IsDirExist(cli.DestDir) {
		return errors.New("invalid file path specified in --dest")
	}
	return nil
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	source, _ := cmd.Flags().GetString("source")
	destination, _ := cmd.Flags().GetString("dest")
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	filenames, _ := cmd.Flags().GetStringSlice("only")

	input := CliInput{
		SourceDir:   source,
		DestDir:     destination,
		IsOverwrite: overwrite,
		Includes:    filenames,
	}

	return input
}
