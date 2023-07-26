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
		var last int
		for i, item := range hunk.ListItems() {
			last = item.Line()
			if i == 0 {
				fmt.Println(ren.getLine(last - 2))
			}
			if item.Added() {
				fmt.Println(color.GreenString("+ " + item.Text()))
			} else {
				fmt.Println(color.RedString("- " + item.Text()))
			}
		}
		fmt.Println(ren.getLine(last))
	}
	// ren.renderSameLine(-1)
	fmt.Printf("\n")
}

func (ren *ContextualRenderer) getLine(line int) string {
	if len(ren.lines) > line {
		return ren.lines[line]
	}
	return ""
}