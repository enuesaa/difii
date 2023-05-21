package diff

import (
	"fmt"
	"strings"
)

type Diffs struct {
	items []string
}
func NewDiffs() *Diffs {
	return &Diffs{
		items: make([]string, 0),
	}
}

func (diffs *Diffs) Add(text string) {
	diffs.items = append(diffs.items, fmt.Sprintf("+ %s", text))
}

func (diffs *Diffs) Remove(text string) {
	diffs.items = append(diffs.items, fmt.Sprintf("- %s", text))	
}

func (diffs *Diffs) ListItems() []string {
	return diffs.items
}

func (diffs *Diffs) Render() string {
	return strings.Join(diffs.items, "\n") + "\n"
}
