package commands

import (
	"github.com/spf13/cobra"
)

func CreateCli() *cobra.Command {
	var cli = createRootCmd()
	cli.AddCommand(createInspectCmd())
	cli.AddCommand(createImportCmd())

	// global options
	cli.PersistentFlags().String("source", "", "Source dir.")
	cli.PersistentFlags().String("dest", "", "Destination dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	cli.PersistentFlags().Bool("no-interactive", false, "Disable interactive prompt.")

	// disable default command or flags.
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.CompletionOptions.DisableDefaultCmd = true

	return cli
}
