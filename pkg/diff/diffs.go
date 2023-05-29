package diff

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

type Diffs struct {
	items []string
}
func NewDiffs() *Diffs {
	return &Diffs{
		items: make([]string, 0),
	}
}

// todo change name to added
func (diffs *Diffs) Add(value Value) {
	diffs.items = append(diffs.items, fmt.Sprintf("+ %s", value.Text()))
}

// todo change name to deleted
func (diffs *Diffs) Remove(value Value) {
	diffs.items = append(diffs.items, fmt.Sprintf("- %s", value.Text()))	
}

func (diffs *Diffs) ListItems() []string {
	return diffs.items
}

func (diffs *Diffs) Render() string {
	return strings.Join(diffs.items, "\n") + "\n"
}

func (diffs *Diffs) RenderWithColor() string {
	ret := ""
	for _, item := range diffs.items {
		if strings.HasPrefix(item, "+") {
			ret += color.GreenString(item + "\n")
		} else {
			ret += color.RedString(item + "\n")
		}
	}
	return ret
}

func (diffs *Diffs) Summary() string {
	add := 0
	remove := 0
	for _, item := range diffs.items {
		if strings.HasPrefix(item, "+") {
			add += 1
		} else {
			remove += 1
		}
	}
	return fmt.Sprintf("+%d -%d", add, remove)
}
