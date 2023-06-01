package main

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/prompt"
	"github.com/spf13/cobra"
)

func main() {
	var command = &cobra.Command{
		Use:  "difii",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := cli.ParseArgs(cmd, args)
			if !input.IsSourceDirSelected() {
				input.SourceDir = prompt.SelectSourceDir()
			}
			if !input.IsDestinationDirSelected() {
				input.DestinationDir = prompt.SelectDestinationDir()
			}
			if err := input.Validate(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return;
			}
			cli.Summary(input)
			cli.Diff(input)
		},
	}

	command.Flags().String("source", "", "Source dir.")
	command.Flags().String("dest", "", "Destination dir.")
	command.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	command.Flags().Bool("overwrite", false, "Overwrite destination file with source file.")
	command.Flags().StringSlice("ignore", make([]string, 0), "Filename to ignore. By default, .git dir is ignored. You can override this behavior.")

	command.Execute()
}
