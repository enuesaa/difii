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
			if !input.IsDestDirSelected() {
				input.DestDir = prompt.SelectDestinationDir()
			}
			if err := input.Validate(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			cli.Summary(input)
			cli.Diff(input)
		},
	}

	command.Flags().String("source", "", "Source dir.")
	command.Flags().String("dest", "", "Destination dir.")
	command.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	command.Flags().Bool("include-git", false, "By default, .git directory is ignored. If you pass this option, you can also diff git directory.")
	command.Flags().Bool("no-interactive", false, "Disable interactive prompt.")
	command.Flags().Bool("overwrite", false, "Overwrite destination file with source file.")

	command.Execute()
}
