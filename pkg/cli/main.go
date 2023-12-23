package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/spf13/cobra"
)

func CreateCli(fsio repo.FsioInterface) *cobra.Command {
	var cli = &cobra.Command{
		Use:     "difii <dir1> <dir2>",
		Short:   "A CLI tool to diff 2 folders.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.0.12",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoFlags() {
				cmd.Help()
				return
			}

			if input.Interactive && !input.IsWorkDirSelected() {
				input.WorkDir = fsio.SelectDir("dir1: ")
			}
			if input.Interactive && !input.IsCompareDirSelected() {
				input.CompareDir = fsio.SelectDir("dir2: ")
			}
			if err := input.Validate(fsio); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}

			switch input.Task {
			case TaskInspect:
				inspectSrv := InspectService{}
				inspectSrv.Render(fsio, input)
			case TaskSummary:
				summarySrv := SummaryService{}
				summarySrv.Plan(fsio, input)
				summarySrv.Render(fsio, input)
			}
		},
	}

	// operations
	cli.Flags().Bool("inspect", false, "Inspect diffs.")
	cli.Flags().StringSlice("only", make([]string, 0), "Specify filename to compare.")
	cli.Flags().BoolP("interactive", "i", false, "Use interactive prompt.")
	cli.Flags().Bool("verbose", false, "Show additional messages.")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.CompletionOptions.DisableDefaultCmd = true
	cli.SilenceUsage = true
	cli.Flags().BoolP("help", "", false, "Show help messages.")
	cli.Flags().BoolP("version", "", false, "Show version information.")
	cli.SetHelpTemplate(`{{.Short}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableFlags}}

Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`)

	return cli
}
