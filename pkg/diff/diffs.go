package diff

import (
	"fmt"
	"github.com/fatih/color"
)

type Diffs struct {
	items []Diffline
}
func NewDiffs() *Diffs {
	return &Diffs{
		items: make([]Diffline, 0),
	}
}

// todo change name to added
func (diffs *Diffs) Add(value Value) {
	diffs.items = append(diffs.items, *NewDiffline(value, true))
}

// todo change name to deleted
func (diffs *Diffs) Remove(value Value) {
	diffs.items = append(diffs.items,  *NewDiffline(value, false))	
}

func (diffs *Diffs) ListItems() []Diffline {
	return diffs.items
}

func (diffs *Diffs) Render() string {
	ret := ""
	for _, item := range diffs.items {
		if item.Added() {
			ret += "+ " + item.Text() + "\n"
		} else {
			ret += "- " + item.Text() + "\n"
		}
	}
	return ret
}

func (diffs *Diffs) RenderWithColor() string {
	ret := ""
	for _, item := range diffs.items {
		if item.Added() {
			ret += color.GreenString("+ " + item.Text() + "\n")
		} else {
			ret += color.RedString("- " + item.Text() + "\n")
		}
	}
	return ret
}

func (diffs *Diffs) Summary() string {
	add := 0
	remove := 0
	for _, item := range diffs.items {
		if item.Added() {
			add += 1
		} else {
			remove += 1
		}
	}
	return fmt.Sprintf("+%d -%d", add, remove)
}
