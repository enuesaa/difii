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

			// summary というより terraform のように diffs を逐次的に出力する
			cli.Summaryline(input)
		},
	}

	return cmd
}
