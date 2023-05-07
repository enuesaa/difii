package cmd

import (
	"fmt"
	"os"
	"io"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/manifoldco/promptui"
	"github.com/enuesaa/difii/pkg/prompt"
)

var Command = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var fromfilepath string

		// choose source dir, not file.
		if len(args) == 0 {
			fmt.Println("Select source dir")
			currentdir, _ := os.Getwd()
			fromfilepath = prompt.ChooseFile(currentdir)
		} else {
			fromfilepath = args[0]
		}

		// choose destination dir, not file.
		var tofilepath string
		if len(args) == 0 {
			fmt.Println("Select destination dir")
			currentdir, _ := os.Getwd()
			tofilepath = prompt.ChooseFile(currentdir)
		} else {
			tofilepath = args[1]
		}
		fmt.Println(fromfilepath)
		fmt.Println(tofilepath)

		// list files
		// run diff here.
		// show diff tables

		fmt.Printf("\nDo you overwrite `%s` to `%s`\n", fromfilepath, tofilepath)
		prompt := promptui.Select{
			Label: "",
			Items: []string{
				"stay",
				"overwrite",
			},
		}
		
		_, result, _ := prompt.Run()
		if result == "overwrite" {
			fromFile, _ := os.Open(fromfilepath)
			defer fromFile.Close()
			toFile, _ := os.Create(tofilepath)
			defer toFile.Close()
			_, _ = io.Copy(toFile, fromFile)
		}

		dmp := diffmatchpatch.New()
		_ = dmp.DiffMain("あ", "あか", false)
		// fmt.Println(dmp.DiffPrettyText(diffs))
	},
}
