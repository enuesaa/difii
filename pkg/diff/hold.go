package diff

import (
	"golang.org/x/exp/slices"
)

// dest基準. dest の文字列を一旦 hold し source との共通文字列が見つかり次第 差分をpushする
type Holder struct {
	dest []string
	source []string
	diffs Diffs
}

func NewHolder() *Holder {
	return &Holder{
		dest: make([]string, 0),
		source: make([]string, 0),
		diffs: *NewDiffs(),
	}
}

func (holder *Holder) HoldDest(text string) {
	holder.dest = append(holder.dest, text)
}

func (holder *Holder) HoldSource(text string) {
	holder.source = append(holder.source, text)
}

func (holder *Holder) GetHoldDestIndex(text string) int {
	// see https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
	i := slices.IndexFunc(holder.dest, func(value string) bool {
		return value == text
	})
	return i
}

func (holder *Holder) Flush() {
	for i, text := range holder.source {
		matched := holder.GetHoldDestIndex(text)
		if matched == -1 {
			continue;
		}

		holder.rebaseSource(i)
		holder.rebaseDest(matched)
	}
}

func (holder *Holder) FlushRest() {
	for _, value := range holder.source {
		holder.markAdd(value)
	}
	for _, value := range holder.dest {
		holder.markRemove(value)
	}
}

func (holder *Holder) rebaseSource(baseIndex int) {
	nextHolds := make([]string, 0)
	for i, text := range holder.source {
		if i < baseIndex {
			// source text does not exist in dest. so mark this text as add-diff
			holder.markAdd(text)
			continue;
		}
		if i == baseIndex {
			continue;
		}
		nextHolds = append(nextHolds, text)
	}
	holder.source = nextHolds
}

func (holder *Holder) rebaseDest(baseIndex int) {
	nextHolds := make([]string, 0)
	for i, value := range holder.dest {
		if i < baseIndex {
			// dest text does not exist in source. so mark this text as remove-diff
			holder.markRemove(value)
			continue;
		}
		if i == baseIndex {
			// source text exists in dest. so do nothing.
			continue;
		}
		nextHolds = append(nextHolds, value)
	}
	holder.dest = nextHolds
}

func (holder *Holder) markAdd(text string) {
	holder.diffs.Add(text)
}

func (holder *Holder) markRemove(text string) {
	holder.diffs.Remove(text)
}

func (holder *Holder) GetDiffs() *Diffs {
	return &holder.diffs
}
