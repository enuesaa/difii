package commands

import (
	"github.com/spf13/cobra"
)

func CreateCli() *cobra.Command {
	var cli = createRootCmd()
	cli.AddCommand(createImportCmd())

	// global options
	cli.Flags().String("source", "", "Source dir.")
	cli.Flags().String("dest", "", "Destination dir.")
	cli.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	cli.Flags().Bool("no-interactive", false, "Disable interactive prompt.")

	return cli
}