package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli/prompt"
	"github.com/spf13/cobra"
)

func CreateCli() *cobra.Command {
	var cli = &cobra.Command{
		Use:     "difii",
		Short:   "A CLI tool to inspect diffs interactively.",
		Args:    cobra.MinimumNArgs(0),
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)
			if input.HasNoFlags() {
				cmd.Help()
				return
			}

			// global options
			if input.Interactive && !input.IsCompareDirSelected() {
				input.CompareDir = prompt.SelectCompareDir()
			}
			if !input.IsWorkDirSelected() {
				input.WorkDir = "."
			}
			if err := input.Validate(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("\n")

			// operations
			if input.Summary {
				ShowDiffsSummary(input)
			}
			if input.Inspect {
				ShowDiffs(input)
			}
			if input.Apply {
				Apply(input)
			}
		},
	}

	// operations
	cli.Flags().Bool("summary", false, "Show diffs summary.")
	cli.Flags().Bool("inspect", false, "Inspect diffs.")
	cli.Flags().Bool("apply", false, "Overwrite working files with comparison.")

	// options
	cli.PersistentFlags().String("compare", "", "Compare dir.")
	cli.PersistentFlags().String("workdir", "", "Working dir. Default value is current dir.")
	cli.PersistentFlags().StringSlice("only", make([]string, 0), "Filename to compare")
	cli.PersistentFlags().BoolP("interactive", "i", false, "Start interactive prompt.")
	// cli.PersistentFlags().BoolP("auto-approve", "", false, "Auto approve.")
	// cli.PersistentFlags().String("report-file", "", "report filename.")

	// disable default behavior
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.CompletionOptions.DisableDefaultCmd = true
	cli.PersistentFlags().SortFlags = false

	// see https://github.com/spf13/cobra/issues/1328
	cli.InitDefaultHelpFlag()
	cli.Flags().MarkHidden("help")
	cli.PersistentFlags().BoolP("help", "", false, "help")

	cli.InitDefaultVersionFlag()
	cli.Flags().MarkHidden("version")
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
