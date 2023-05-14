package main

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/cli"
	// "github.com/enuesaa/difii/pkg/diff"
	"github.com/enuesaa/difii/pkg/files"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use: "difii",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		input := cli.ParseArgs(cmd, args)
		if !input.IsSourceDirSelected() {
			fmt.Println("source dir not selected")
			cli.DoPrompt()
			// input.SourceDir = cli.ChooseSourceDir()
			return;
		}
		if !input.IsDestinationDirSelected() {
			fmt.Println("destination dir not selected")
			// input.SourceDir = cli.ChooseDestinationDir()
			return;
		}
		fmt.Printf("source dir: %s \n", input.SourceDir)
		fmt.Printf("destination dir: %s \n", input.DestinationDir)
		fmt.Println("")

		sourcefiles := files.ListFilesRecursively(input.SourceDir)
		files.ReadStreamWithDiff(input.SourceDir, input.DestinationDir, sourcefiles[0])

		// for _, filename := range sourcefiles {
		// 	sourcefile := files.Read(input.SourceDir, filename)
		// 	destinationfile := files.Read(input.DestinationDir, filename)
		// 	diff.Diff(sourcefile, destinationfile)
		// 	// run diff here.
		// }

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
	RootCommand.Flags().String("source", "", "Source directory.")
	RootCommand.Flags().String("destination", "", "Destination directory.")
	RootCommand.Flags().StringSlice("only", make([]string, 0), "Filename to compare")
	RootCommand.Flags().Bool("overwrite", false, "Overwrite destination file with source file.")
	RootCommand.Flags().StringSlice("ignore", make([]string, 0), "Filename to ignore. By default, .git dir is ignored. You can override this behavior.")
	RootCommand.Execute()
}
