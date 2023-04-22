package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var RootCmd = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(1),
	Run: rootCmdHandler,
}

func rootCmdHandler(cmd *cobra.Command, args []string)  {
	dirname := args[0]
	files, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Printf("%s", file)
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain("あ", "あか", false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
