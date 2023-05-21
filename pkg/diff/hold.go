package diff

import (
	"golang.org/x/exp/slices"
)

// dest基準. dest の文字列を一旦 hold し source との共通文字列が見つかり次第 差分をpushする
type Holder struct {
	holds []string
	diffs Diffs
}

func NewHolder() *Holder {
	return &Holder{ holds: make([]string, 0), diffs: *NewDiffs() }
}

func (holder *Holder) Hold(text string) {
	holder.holds = append(holder.holds, text)
}

func (holder *Holder) GetHoldIndex(text string) int {
	// see https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
	i := slices.IndexFunc(holder.holds, func(value string) bool {
		return value == text
	})
	return i
}

func (holder *Holder) Flush(text string) {
	matched := holder.GetHoldIndex(text)
	if matched == -1 {
		// source text does not exist in dest. so mark this text as add-diff
		holder.markAdd(text)
		return;
	}

	nextHolds := make([]string, 0)
	for i, value := range holder.holds {
		if i < matched {
			// dest text does not exist in source. so mark this text as remove-diff
			holder.markRemove(value)
			continue;
		}
		if i == matched {
			// source text exists in dest. so do nothing.
			continue;
		}
		nextHolds = append(nextHolds, value)
	}
	holder.holds = nextHolds
}

func (holder *Holder) FlushRest() {
	for _, value := range holder.holds {
		holder.markRemove(value)
	}
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
