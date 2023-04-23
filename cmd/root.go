package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var RootCmd = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(2),
	Run: rootCmdHandler,
}

func rootCmdHandler(cmd *cobra.Command, args []string)  {
	fromdir := args[0]
	fromfiles, err := os.ReadDir(fromdir)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	todir := args[1]
	tofiles, err := os.ReadDir(todir)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
	fmt.Printf("%+v", tofiles)

	for _, file := range fromfiles {
		if file.IsDir() {
			fmt.Printf("dir: %s\n", file.Name())
		} else {
			info, _ := file.Info()
			fmt.Printf("file: %s %d\n", file.Name(), info.Size())
		}
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain("あ", "あか", false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
