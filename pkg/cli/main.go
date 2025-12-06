package cli

import (
	"log"

	"github.com/enuesaa/difii/pkg/repository"
	"github.com/spf13/cobra"
)

func New(repos *repository.Repos) *cobra.Command {
	app := &cobra.Command{
		Use:     "difii <dir1> <dir2>",
		Short:   "A CLI tool to diff 2 folders interactively.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.0.16",
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
	app.Flags().Bool("inspect", false, "Inspect diffs.")
	app.Flags().StringSlice("only", make([]string, 0), "Specify filename to compare.")
	app.Flags().BoolP("interactive", "i", false, "Use interactive prompt.")

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.Flags().BoolP("help", "", false, "Show help messages.")
	app.Flags().BoolP("version", "", false, "Show version information.")
	app.SetHelpTemplate(`{{.Short}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableFlags}}

Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`)

	return app
}
