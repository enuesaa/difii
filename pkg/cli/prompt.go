package cli

import (
	"os"
	"strings"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func ChooseSourceDir() string {
	current, _  := os.Getwd()
	return chooseFile(current)
}

func ChooseDestinationDir() string {
	current, _  := os.Getwd()
	return chooseFile(current)
}

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