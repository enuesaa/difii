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
	Inspect     bool
	Apply       bool
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
	return !cli.Inspect && !cli.Apply
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
	if len(args) > 0 {
		compareDir = args[0]
	}
	workDir, _ := cmd.Flags().GetString("workdir")
	includes, _ := cmd.Flags().GetStringSlice("only")
	interactive, _ := cmd.Flags().GetBool("interactive")

	inspect, _ := cmd.Flags().GetBool("inspect")
	apply, _ := cmd.Flags().GetBool("apply")

	input := CliInput{
		CompareDir:  compareDir,
		WorkDir:     workDir,
		Includes:    includes,
		Interactive: interactive,
		Inspect:     inspect,
		Apply:       apply,
	}

	return input
}
