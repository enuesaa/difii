package cmd

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/manifoldco/promptui"
)

var Command = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fromdir := args[0]
		todir := args[1]

		fromfiles, err := os.ReadDir(fromdir)
		if err != nil {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}

		_, err = os.ReadDir(todir)
		if err != nil {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}

		for _, file := range fromfiles {
			if !file.IsDir() {
				fromPath := filepath.Join(fromdir, file.Name())
				toPath := filepath.Join(todir, file.Name())
				fmt.Printf("\nDo you overwrite `%s` to `%s`\n", fromPath, toPath)

				prompt := promptui.Select{
					Label: "",
					Items: []string{
						"stay",
						"overwrite",
					},
				}
				
				_, result, _ := prompt.Run()
				if result == "overwrite" {
					fromFile, _ := os.Open(fromPath)
					defer fromFile.Close()
					toFile, _ := os.Create(toPath)
					defer toFile.Close()
					_, _ = io.Copy(toFile, fromFile)
				}
			}
		}

		dmp := diffmatchpatch.New()
		_ = dmp.DiffMain("あ", "あか", false)
		// fmt.Println(dmp.DiffPrettyText(diffs))
	},
}
