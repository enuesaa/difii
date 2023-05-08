package main

import (
	// "os"
	// "io"
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/sergi/go-diff/diffmatchpatch"
	// "github.com/manifoldco/promptui"
	"github.com/enuesaa/difii/pkg/cli"
)

var RootCommand = &cobra.Command{
	Use: "difii",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		input := cli.ParseArgs(cmd, args)
		if !input.IsSourceDirSelected() {
			input.SourceDir = cli.ChooseSourceDir()
		}
		if !input.IsDestinationDirSelected() {
			input.SourceDir = cli.ChooseDestinationDir()
		}
		fmt.Printf("%+v", input)

		// // list files
		// // run diff here.
		// // show diff tables

		// fmt.Printf("\nDo you overwrite `%s` to `%s`\n", fromfilepath, tofilepath)
		// prompt := promptui.Select{
		// 	Label: "",
		// 	Items: []string{
		// 		"stay",
		// 		"overwrite",
		// 	},
		// }
		
		// _, result, _ := prompt.Run()
		// if result == "overwrite" {
		// 	fromFile, _ := os.Open(fromfilepath)
		// 	defer fromFile.Close()
		// 	toFile, _ := os.Create(tofilepath)
		// 	defer toFile.Close()
		// 	_, _ = io.Copy(toFile, fromFile)
		// }

		// dmp := diffmatchpatch.New()
		// _ = dmp.DiffMain("あ", "あか", false)
		// fmt.Println(dmp.DiffPrettyText(diffs))
	},
}

func main() {
	RootCommand.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	RootCommand.Flags().Bool("overwrite", false, "Overwrite destination file with source file.")
	RootCommand.Execute()
}
