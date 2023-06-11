package commands

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/prompt"
	"github.com/enuesaa/difii/pkg/cli"
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
			cli.Summaryline(input)

			fmt.Printf("\n")
			fmt.Printf("To inspect diffs, please run command below.\n")
			fmt.Printf("  difii --source %s --dest %s inspect\n", input.SourceDir, input.DestDir)
			fmt.Printf("\n")
		},
	}

	return cmd
}
