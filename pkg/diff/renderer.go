package diff

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

// experimental
type Renderer interface {
	Render() // できればここで値を返したい
}

// これをデフォルトにして、色とかは Option にしたい
type HunkedRenderer struct {
	diffs Diffs
}
func NewHunkedRenderer(diffs Diffs) *HunkedRenderer {
	return &HunkedRenderer{
		diffs,
	}
}

func (ren *HunkedRenderer) Render() {
	hunked := make([]int, 0)

	for _, item := range ren.diffs.items {
		line := item.Line()
		if len(hunked) == 0 {
			ren.renderLine(item)
			hunked = append(hunked, item.Line())
			continue
		}

		if slices.Contains(hunked, line) {
			ren.renderLine(item)
			continue
		}

		if slices.Contains(hunked, line - 1) || slices.Contains(hunked, line + 1) {
			ren.renderLine(item)
			hunked = append(hunked, item.Line())
			continue
		}

		prompt.Input("Do you overwrite ? [Y/n] ", func (in prompt.Document) []prompt.Suggest {
			return make([]prompt.Suggest, 0)
		})

		fmt.Println("")
		ren.renderLine(item)
		hunked = make([]int, 0)
		hunked = append(hunked, item.Line())
	}
}

func (ren *HunkedRenderer) renderLine(item Diffline) {
	if item.Added() {
		fmt.Println(color.GreenString("+ " + item.Text()))
	} else {
		fmt.Println(color.RedString("- " + item.Text()))
	}
}
