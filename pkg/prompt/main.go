package prompt

import (
	"github.com/c-bata/go-prompt"
)

func SelectSourceDir() string {
	return prompt.Input("Select Source Dir ", selectDir, promptOptions()...)
}

func SelectDestinationDir() string {
	return prompt.Input("Select Destination Dir ", selectDir, promptOptions()...)
}
