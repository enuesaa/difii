package diff

import (
	"slices"
)

type Diffs struct {
	items []Diffline
}

func NewDiffs() *Diffs {
	return &Diffs{
		items: make([]Diffline, 0),
	}
}

func (diffs *Diffs) MarkAdd(value Value) {
	diffs.items = append(diffs.items, *NewDiffline(value.Line(), value.Text(), Added))
}

func (diffs *Diffs) MarkRemove(value Value) {
	diffs.items = append(diffs.items, *NewDiffline(value.Line(), value.Text(), Removed))
}

func (diffs *Diffs) ListItems() []Diffline {
	return diffs.items
}

func (diffs *Diffs) ListHunks() []Hunk {
	hunks := make([]Hunk, 0)
	staging := make([]int, 0)

	hunk := NewHunk()
	for _, item := range diffs.items {
		line := item.Line()
		if len(staging) == 0 {
			hunk.Push(item)
			staging = append(staging, item.Line())
			continue
		}

		if slices.Contains(staging, line) {
			hunk.Push(item)
			continue
		}

		if slices.Contains(staging, line-1) || slices.Contains(staging, line+1) {
			hunk.Push(item)
			staging = append(staging, item.Line())
			continue
		}

		hunks = append(hunks, *hunk)
		hunk = NewHunk()
		hunk.Push(item)
		staging = make([]int, 0)
		staging = append(staging, item.Line())
	}

	// last loop
	if len(staging) > 0 {
		hunks = append(hunks, *hunk)
	}

	return hunks
}

func (diffs *Diffs) CountAdd() int {
	add := 0
	for _, item := range diffs.items {
		if item.Added() {
			add += 1
		}
	}
	return add
}

func (diffs *Diffs) CountRemove() int {
	remove := 0
	for _, item := range diffs.items {
		if !item.Added() {
			remove += 1
		}
	}
	return remove
}
