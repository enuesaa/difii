package cli

import (
	"errors"

	"github.com/enuesaa/difii/pkg/files"
	"github.com/spf13/cobra"
)

type CliInput struct {
	CompareDir  string
	BaseDir     string
	Includes    []string
	Interactive bool
	Summary     bool
	Inspect     bool
	Apply       bool
}

func (cli *CliInput) IsCompareDirSelected() bool {
	return cli.CompareDir != ""
}
func (cli *CliInput) IsBaseDirSelected() bool {
	return cli.BaseDir != ""
}
func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
func (cli *CliInput) HasNoOperationFlags() bool {
	return !cli.Summary && !cli.Inspect && !cli.Apply
}
func (cli *CliInput) HasNoGlobalFlags() bool {
	return !cli.IsCompareDirSelected() && !cli.IsBaseDirSelected() && !cli.IsFileSpecified() && !cli.Interactive
}
func (cli *CliInput) Validate() error {
	if !cli.IsCompareDirSelected() {
		return errors.New("required option --compare is missing")
	}
	if !files.IsDirExist(cli.CompareDir) {
		return errors.New("invalid file path specified in --compare")
	}
	if !files.IsDirExist(cli.BaseDir) {
		return errors.New("invalid file path specified in --base")
	}
	return nil
}

func ParseArgs(cmd *cobra.Command) CliInput {
	compareDir, _ := cmd.Flags().GetString("compare")
	baseDir, _ := cmd.Flags().GetString("base")
	includes, _ := cmd.Flags().GetStringSlice("only")
	interactive, _ := cmd.Flags().GetBool("interactive")

	summary, _ := cmd.Flags().GetBool("summary")
	inspect, _ := cmd.Flags().GetBool("inspect")
	apply, _ := cmd.Flags().GetBool("apply")

	input := CliInput{
		CompareDir:  compareDir,
		BaseDir:     baseDir,
		Includes:    includes,
		Interactive: interactive,
		Summary:     summary,
		Inspect:     inspect,
		Apply:       apply,
	}

	return input
}
