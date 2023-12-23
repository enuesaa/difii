package cli

import (
	"errors"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/spf13/cobra"
)

type CliInput struct {
	CompareDir  string
	WorkDir     string
	Includes    []string
	Interactive bool
	Verbose     bool
	Inspect     bool
}

func (cli *CliInput) IsCompareDirSelected() bool {
	return cli.CompareDir != ""
}
func (cli *CliInput) IsWorkDirSelected() bool {
	return cli.WorkDir != ""
}
func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
func (cli *CliInput) HasNoOperationFlags() bool {
	return !cli.Inspect
}
func (cli *CliInput) HasNoGlobalFlags() bool {
	return !cli.IsCompareDirSelected() && !cli.IsWorkDirSelected() && !cli.IsFileSpecified() && !cli.Interactive
}
func (cli *CliInput) Validate(fsio repo.FsioInterface) error {
	if !cli.IsCompareDirSelected() {
		return errors.New("required argument <compare-dir> is missing")
	}
	if !fsio.IsDirOrFileExist(cli.CompareDir) {
		return errors.New("invalid file path specified in <compare-dir>")
	}
	if !fsio.IsDirOrFileExist(cli.WorkDir) {
		return errors.New("invalid file path specified in --workdir")
	}
	return nil
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	compareDir := ""
	workDir := ""
	if len(args) > 0 {
		workDir = args[0]
	}
	if len(args) > 1 {
		compareDir = args[1]
	}
	includes, _ := cmd.Flags().GetStringSlice("only")
	interactive, _ := cmd.Flags().GetBool("interactive")
	verbose, _ := cmd.Flags().GetBool("verbose")
	inspect, _ := cmd.Flags().GetBool("inspect")

	input := CliInput{
		CompareDir:  compareDir,
		WorkDir:     workDir,
		Includes:    includes,
		Interactive: interactive,
		Verbose:     verbose,
		Inspect:     inspect,
	}

	return input
}
