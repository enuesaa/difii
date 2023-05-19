package main

import (
	"github.com/enuesaa/difii/pkg/cli"
	"github.com/spf13/cobra"
)

func main() {
	var command = &cobra.Command{
		Use:  "difii",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := cli.ParseArgs(cmd, args)
			if !input.IsSourceDirSelected() {
				input.SourceDir = cli.SelectSourceDir()
			}
			// if !input.IsDestinationDirSelected() {
			// 	input.DestinationDir = cli.SelectDestinationDir()
			// }
			// cli.DiffFiles(input.SourceDir, input.DestinationDir)
		},
	}

	command.Flags().String("source", "", "Source directory.")
	command.Flags().String("destination", "", "Destination directory.")
	command.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	command.Flags().Bool("overwrite", false, "Overwrite destination file with source file.")
	command.Flags().StringSlice("ignore", make([]string, 0), "Filename to ignore. By default, .git dir is ignored. You can override this behavior.")

	command.Execute()
}
