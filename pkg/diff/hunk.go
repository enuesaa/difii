package diff

type Hunk struct {
	items []Diffline
}

func NewHunk() *Hunk {
	return &Hunk{
		items: make([]Diffline, 0),
	}
}

func (hunk *Hunk) Push(diffline Diffline) {
	hunk.items = append(hunk.items, diffline)
}

func (hunk *Hunk) ListItems() []Diffline {
	return hunk.items
}
