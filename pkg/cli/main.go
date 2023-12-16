package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/spf13/cobra"
)

func CreateCli(fsio repo.FsioInterface) *cobra.Command {
	var cli = &cobra.Command{
		Use:     "difii <dir1> <dir2>",
		Short:   "A CLI tool to inspect diffs interactively.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.0.12",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoOperationFlags() && input.HasNoGlobalFlags() {
				cmd.Help()
				return
			}

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

			summarySrv := SummaryService{}
			summarySrv.Plan(fsio, input)
			summarySrv.Render(fsio, input)

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
	cli.Flags().StringSlice("only", make([]string, 0), "Specify filename to compare.")
	cli.Flags().BoolP("interactive", "i", false, "Use interactive prompt.")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.CompletionOptions.DisableDefaultCmd = true
	cli.SilenceUsage = true
	cli.PersistentFlags().SortFlags = false
	cli.PersistentFlags().BoolP("help", "", false, "Show help")
	cli.PersistentFlags().BoolP("version", "", false, "Show version")
	cli.SetHelpTemplate(`{{.Short}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableFlags}}

Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`)

	return cli
}
