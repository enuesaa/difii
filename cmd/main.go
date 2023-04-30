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

// todo refactor
// see https://github.com/manifoldco/promptui/blob/master/example_select_test.go
func chooseDir(dir string) string {
	files, _ := os.ReadDir(dir)
	filenames := make([]string, 0)
	filenames = append(filenames, "[ok]", "../")
	for _, file := range files {
		if file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	prompt := promptui.Select{
		Label: "",
		Items: filenames,
		Size: len(filenames),
	}
	_, result, _ := prompt.Run()
	if result == "../" {
		fmt.Println(dir)
		result = chooseDir(filepath.Dir(dir))
	} else if result == "[ok]" {
		result = dir
	} else {
		result = chooseDir(dir + "/" + result)
	}
	return result
}

var Command = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var fromdir string
		if len(args) == 0 {
			fmt.Println("Select from dir")
			currentdir, _ := os.Getwd()
			fromdir = chooseDir(currentdir)
		} else {
			fromdir = args[0]
		}

		var todir string
		if len(args) == 0 {
			fmt.Println("Select to dir")
			currentdir, _ := os.Getwd()
			todir = chooseDir(currentdir)
		} else {
			todir = args[1]
		}
		fmt.Println(fromdir)
		fmt.Println(todir)

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
