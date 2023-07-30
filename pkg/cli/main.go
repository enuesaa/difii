package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "difii",
		Short: "A CLI tool to inspect diffs interactively.",
		Args: cobra.MinimumNArgs(0),
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoFlags() {
				cmd.Help()
				return
			}

			if !input.IsCompareDirSelected() {
				input.CompareDir = prompt.SelectCompareDir()
			}
			if !input.IsWorkDirSelected() {
				input.WorkDir = "."
			}
			if err := input.Validate(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("\n")

			if input.Summary {
				ShowDiffsSummary(input)
			}
			if input.Inspect {
				ShowDiffs(input)
			} else {
				RecommendInspectFlag(input)
			}

			if input.Apply {
				prompt.ConfirmToApply()
				// todo
			}
			// else {
			// 	RecommendApplyFlag(input)
			// }
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	var cli = createRootCmd()

	// options
	cli.PersistentFlags().String("compare", "", "Compare dir.")
	cli.PersistentFlags().String("workdir", "", "Working dir. Default value is current dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	cli.PersistentFlags().Bool("inspect", false, "Inspect diffs.")
	cli.PersistentFlags().Bool("apply", false, "Overwrite working files with comparison.")
	cli.PersistentFlags().Bool("summary", false, "Show diffs summary.")
	// cli.PersistentFlags().Bool("report", false, "Report file name.")
	// cli.PersistentFlags().Bool("no-interactive", false, "Disable interactive prompt.")

	// disable default behavior.
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true
	// cli.SetUsageTemplate(getUsageTemplate())

	return cli
}