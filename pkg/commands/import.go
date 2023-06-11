package commands

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/prompt"
	"github.com/enuesaa/difii/pkg/commands/common"
	"github.com/spf13/cobra"
)

func createImportCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "difii import",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := common.ParseArgs(cmd, args)
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
		},
	}

	return cmd
}
