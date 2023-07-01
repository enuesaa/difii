package cli

import (
	"errors"

	"github.com/enuesaa/difii/pkg/files"
	"github.com/spf13/cobra"
)

type CliInput struct {
	CompareDir   string
	WorkDir     string
	IsOverwrite bool
	Includes    []string
}

func (cli *CliInput) IsFileSpecified() bool {
	return len(cli.Includes) > 0
}
func (cli *CliInput) IsCompareDirSelected() bool {
	return cli.CompareDir != ""
}
func (cli *CliInput) IsWorkDirSelected() bool {
	return cli.WorkDir != ""
}
func (cli *CliInput) Validate() error {
	if !cli.IsCompareDirSelected() {
		return errors.New("required option --compare is missing")
	}
	if !files.IsDirExist(cli.CompareDir) {
		return errors.New("invalid file path specified in --compare")
	}
	if !files.IsDirExist(cli.WorkDir) {
		return errors.New("invalid file path specified in --workdir")
	}
	return nil
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	compareDir, _ := cmd.Flags().GetString("compare")
	workDir, _ := cmd.Flags().GetString("workdir")
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	filenames, _ := cmd.Flags().GetStringSlice("only")

	input := CliInput{
		CompareDir:  compareDir,
		WorkDir:     workDir,
		IsOverwrite: overwrite,
		Includes:    filenames,
	}

	return input
}
