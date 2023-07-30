package cli

import (
	// "fmt"

	// "github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "difii",
		Short: "A CLI tool to inspect diffs interactively.",
		Args: cobra.MinimumNArgs(0),
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			// input := cli.ParseArgs(cmd, args)
			// if !input.IsCompareDirSelected() {
			// 	input.CompareDir = prompt.SelectCompareDir()
			// }
			// if !input.IsWorkDirSelected() {
			// 	input.WorkDir = "."
			// }
			// if err := input.Validate(); err != nil {
			// 	fmt.Printf("Error: %s\n", err.Error())
			// 	return
			// }

			// fmt.Printf("\n")
			// cli.ShowDiffsSummary(input)
			// cli.RecommendInspectCmd(input)

			// inspect
			// cli.ShowDiffs(input)

		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	var cli = createRootCmd()

	// global options
	cli.PersistentFlags().String("compare", "", "Compare dir.")
	cli.PersistentFlags().String("workdir", "", "Working dir. Default value is current dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	// cli.Flags().Bool("no-interactive", false, "Disable interactive prompt.")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true
	// cli.SetUsageTemplate(getUsageTemplate())

	return cli
}