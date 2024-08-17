package cli

import (
	"log"

	"github.com/enuesaa/difii/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateCli(repos repository.Repos) *cobra.Command {
	var cli = &cobra.Command{
		Use:     "difii <dir1> <dir2>",
		Short:   "A CLI tool to diff 2 folders interactively.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.0.15",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoFlags() {
				cmd.Help()
				return
			}

			if input.Interactive && !input.IsWorkDirSelected() {
				input.WorkDir = repos.Fsio.SelectDir("dir1: ")
			}
			if input.Interactive && !input.IsCompareDirSelected() {
				input.CompareDir = repos.Fsio.SelectDir("dir2: ")
			}
			if err := input.Validate(repos.Fsio); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}

			switch input.Task {
			case TaskInspect:
				inspectSrv := NewInspectService(repos)
				inspectSrv.Render(input)
			case TaskSummary:
				summarySrv := NewSummaryService(repos)
				summarySrv.Plan(input)
				summarySrv.Render(input)
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
