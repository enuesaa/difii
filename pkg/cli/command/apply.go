package command

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/spf13/cobra"
)

func createApplyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "apply",
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

			// todo apply
		},
	}

	return cmd
}
