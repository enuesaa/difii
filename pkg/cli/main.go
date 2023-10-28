package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/spf13/cobra"
)

func CreateCli(fsio repo.FsioInterface) *cobra.Command {
	var cli = &cobra.Command{
		Use:     "difii <compare-dir>",
		Short:   "A CLI tool to inspect diffs interactively.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.0.10",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoOperationFlags() && input.HasNoGlobalFlags() {
				cmd.Help()
				return
			}

			// options
			if input.Interactive && !input.IsCompareDirSelected() {
				input.CompareDir = fsio.SelectCompareDir()
			}
			if !input.IsWorkDirSelected() {
				input.WorkDir = "."
			}
			if err := input.Validate(fsio); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			Plan(fsio, input)

			summarySrv := SummaryService{}
			if !input.Interactive || summarySrv.Confirm(fsio) {
				summarySrv.Render(fsio, input)
			}

			inspectSrv := InspectService{}
			if input.Interactive {
				input.Inspect = inspectSrv.Confirm(fsio)
			}
			if input.Inspect {
				inspectSrv.Render(fsio, input)
			}
		},
	}

	// operations
	cli.Flags().Bool("inspect", false, "Inspect diffs.")

	// options
	cli.PersistentFlags().String("workdir", "", "Working dir. Default value is current dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	cli.PersistentFlags().BoolP("no-interactive", "", false, "Do not use interactive prompt.")

	// disable default behavior
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.CompletionOptions.DisableDefaultCmd = true
	// see https://github.com/spf13/cobra/issues/340
	cli.SilenceUsage = true
	cli.PersistentFlags().SortFlags = false
	cli.PersistentFlags().BoolP("help", "", false, "help")
	cli.PersistentFlags().BoolP("version", "", false, "version")
	cli.SetHelpTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableFlags}}

Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailablePersistentFlags}}

Global Flags:
{{.PersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`)

	return cli
}
