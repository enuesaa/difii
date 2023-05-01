package cmd

import (
	"fmt"
	"os"
	"io"
	"strings"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/manifoldco/promptui"
)

// todo refactor
// see https://github.com/manifoldco/promptui/blob/master/example_select_test.go
func chooseFile(dir string) string {
	var choosed string

	files, _ := os.ReadDir(dir)
	filenames := make([]string, 0)
	filenames = append(filenames, "../")
	for _, file := range files {
		if file.IsDir() {
			filenames = append(filenames, file.Name() + "/")
		} else {
			filenames = append(filenames, file.Name())
		}
	}

	prompt := promptui.Select{
		Label: "",
		Items: filenames,
		Size: len(filenames),
	}
	_, result, _ := prompt.Run()
	if strings.HasSuffix(result, "/") {
		if result == "../" {
			choosed = chooseFile(filepath.Dir(dir))
		} else {
			choosed = chooseFile(filepath.Join(dir, result))
		}
	} else {
		choosed = filepath.Join(dir, result)
	}

	return choosed
}

var Command = &cobra.Command{
	Use:   "difii",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var fromfilepath string
		if len(args) == 0 {
			fmt.Println("Select source file")
			currentdir, _ := os.Getwd()
			fromfilepath = chooseFile(currentdir)
		} else {
			fromfilepath = args[0]
		}

		var tofilepath string
		if len(args) == 0 {
			fmt.Println("Select destination file")
			currentdir, _ := os.Getwd()
			tofilepath = chooseFile(currentdir)
		} else {
			tofilepath = args[1]
		}
		fmt.Println(fromfilepath)
		fmt.Println(tofilepath)

		if _, err := os.Stat(fromfilepath); err != nil {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}
		if _, err := os.Stat(tofilepath); err != nil {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}

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
