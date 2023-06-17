package command

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "difii",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := cli.ParseArgs(cmd, args)
			if !input.IsSourceDirSelected() {
				input.SourceDir = prompt.SelectSourceDir()
			}
			if !input.IsDestDirSelected() {
				input.DestDir = prompt.SelectDestinationDir()
			}
			if err := input.Validate(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}

			fmt.Printf("\n")
			cli.ShowDiffsSummary(input)
			cli.RecommendInspectCmd(input)
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	var cli = createRootCmd()
	cli.AddCommand(createApplyCmd())
	cli.AddCommand(createInspectCmd())

	// global options
	cli.PersistentFlags().String("source", "", "Source dir.")
	cli.PersistentFlags().String("dest", "", "Destination dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	// cli.Flags().Bool("no-interactive", false, "Disable interactive prompt.")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true
	// cli.SetUsageTemplate(getUsageTemplate())

	return cli
}
