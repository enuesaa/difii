package prompt

import (
	"os"
	"strings"
	"path/filepath"
	"github.com/manifoldco/promptui"
)

// todo refactor
// see https://github.com/manifoldco/promptui/blob/master/example_select_test.go
// see https://qiita.com/sueken/items/87093e5941bfbc09bea8
func ChooseFile(dir string) string {
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
			choosed = ChooseFile(filepath.Dir(dir))
		} else {
			choosed = ChooseFile(filepath.Join(dir, result))
		}
	} else {
		choosed = filepath.Join(dir, result)
	}

	return choosed
}
