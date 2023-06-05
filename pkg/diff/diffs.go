package diff

import "golang.org/x/exp/slices"

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
	diffs.items = append(diffs.items, *NewDiffline(value, false))
}

func (diffs *Diffs) ListItems() []Diffline {
	return diffs.items
}

func (diffs *Diffs) ListHunks() []Hunk {
	hunks := make([]Hunk, 0)
	staging := make([]int, 0)

	list := make([]string, 0)
	list = append(list, "orange")
	list = append(list, "apple")
	list = append(list, "blueberry")
	list = append(list, "cherry")

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

		hunk.Push(item)
		hunks = append(hunks, *hunk)
		hunk = NewHunk()
		staging = make([]int, 0)
		staging = append(staging, item.Line())
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
