package prompt

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/difii/pkg/files"
)

// see https://github.com/c-bata/go-prompt/issues/8
func SelectSourceDir() string {
	for {
		dir := prompt.Input("Source dir (--source): ", selectDir, promptOptions()...)
		if files.IsDirExist(dir) {
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func SelectDestinationDir() string {
	for {
		dir := prompt.Input("Destination dir (--dest): ", selectDir, promptOptions()...)
		if files.IsDirExist(dir) {
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}
