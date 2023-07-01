package prompt

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/difii/pkg/files"
)

// see https://github.com/c-bata/go-prompt/issues/8
func SelectCompareDir() string {
	saveState()
	for {
		dir := prompt.Input("Compare dir (--compare): ", selectDir, promptOptions()...)
		if files.IsDirExist(dir) {
			restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}
