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


// func (ren *HunkedRenderer) Render() {
// 	hunked := make([]int, 0)

// 	for _, item := range ren.diffs.items {
// 		line := item.Line()
// 		if len(hunked) == 0 {
// 			ren.renderLine(item)
// 			hunked = append(hunked, item.Line())
// 			continue
// 		}

// 		if slices.Contains(hunked, line) {
// 			ren.renderLine(item)
// 			continue
// 		}

// 		if slices.Contains(hunked, line - 1) || slices.Contains(hunked, line + 1) {
// 			ren.renderLine(item)
// 			hunked = append(hunked, item.Line())
// 			continue
// 		}

// 		prompt.Input("Do you overwrite ? [Y/n] ", func (in prompt.Document) []prompt.Suggest {
// 			return make([]prompt.Suggest, 0)
// 		})

// 		fmt.Println("")
// 		ren.renderLine(item)
// 		hunked = make([]int, 0)
// 		hunked = append(hunked, item.Line())
// 	}
// }

// func (ren *HunkedRenderer) renderLine(item Diffline) {
// 	if item.Added() {
// 		fmt.Println(color.GreenString("+ " + item.Text()))
// 	} else {
// 		fmt.Println(color.RedString("- " + item.Text()))
// 	}
// }