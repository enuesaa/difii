package render

import (
	"fmt"
	"io"
	"strings"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/fatih/color"
)

type ContextualRenderer struct {
	diffs diff.Diffs
	lines []string
}

func NewContextualRenderer(diffs diff.Diffs, dest io.Reader) *ContextualRenderer {
	raw, _ := io.ReadAll(dest)
	lines := strings.Split(string(raw), "\n")

	return &ContextualRenderer{
		diffs: diffs,
		lines: lines,
	}
}

func (ren *ContextualRenderer) Render() {
	for _, hunk := range ren.diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			if item.Added() {
				fmt.Println(color.GreenString("+ " + item.Text()))
			} else {
				fmt.Println(color.RedString("- " + item.Text()))
			}
		}
	}
	fmt.Printf("\n")
}
