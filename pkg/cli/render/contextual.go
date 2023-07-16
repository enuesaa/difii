package render

import (
	"fmt"
	"bufio"
	"io"

	"github.com/enuesaa/difii/pkg/diff"
	"github.com/fatih/color"
)

type ContextualRenderer struct {
	diffs diff.Diffs
	dest  bufio.Scanner
	last int
}

func NewContextualRenderer(diffs diff.Diffs, dest io.Reader) *ContextualRenderer {
	return &ContextualRenderer{
		diffs: diffs,
		dest: *bufio.NewScanner(dest),
		last: 0,
	}
}

func (ren *ContextualRenderer) Render() {
	for _, hunk := range ren.diffs.ListHunks() {
		for _, item := range hunk.ListItems() {
			ren.renderSameLine(item.Line())
			if item.Added() {
				fmt.Println(color.GreenString("+ " + item.Text()))
			} else {
				fmt.Println(color.RedString("- " + item.Text()))
			}
		}
	}
	// ren.renderSameLine(-1)
	fmt.Printf("\n")
}

// Pattern:
//   last:0, next:3   .. 一番初めのhunk
//   last:6, next:10  .. 次のhunk
//   last:11, next:-1 .. 最後
func (ren *ContextualRenderer) renderSameLine(current int) {
	if ren.last > current {
		return;
	}
	// -1 
	if ren.last == current - 1 {
		ren.dest.Scan()
		ren.last += 1
		return;
	}
	for i := ren.last; i < current - 1; i++ {
		ren.dest.Scan()
		if i >= current - 2 {
			fmt.Println(ren.dest.Text())
		}
	}
	ren.last = current
}
