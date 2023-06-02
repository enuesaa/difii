package diff

import (
	"golang.org/x/exp/slices"
)

type Holder struct {
	dest   []Value
	source []Value
	diffs  Diffs
}

func NewHolder() *Holder {
	return &Holder{
		dest:   make([]Value, 0),
		source: make([]Value, 0),
		diffs:  *NewDiffs(),
	}
}

func (holder *Holder) HoldDest(value Value) {
	holder.dest = append(holder.dest, value)
}

func (holder *Holder) HoldSource(value Value) {
	holder.source = append(holder.source, value)
}

func (holder *Holder) GetHoldDestIndex(source Value) int {
	// see https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
	i := slices.IndexFunc(holder.dest, func(value Value) bool {
		return value.Text() == source.Text()
	})
	return i
}

func (holder *Holder) Flush() {
	for i, sourceValue := range holder.source {
		matched := holder.GetHoldDestIndex(sourceValue)
		if matched == -1 {
			continue
		}

		holder.rebaseDest(matched)
		holder.rebaseSource(i)
	}
}

func (holder *Holder) FlushRest() {
	for _, value := range holder.dest {
		holder.markRemove(value)
	}
	for _, value := range holder.source {
		holder.markAdd(value)
	}
}

func (holder *Holder) rebaseSource(baseIndex int) {
	nextHolds := make([]Value, 0)
	for i, value := range holder.source {
		if i < baseIndex {
			// source text does not exist in dest. so mark this text as add-diff
			holder.markAdd(value)
			continue
		}
		if i == baseIndex {
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	holder.source = nextHolds
}

func (holder *Holder) rebaseDest(baseIndex int) {
	nextHolds := make([]Value, 0)
	for i, value := range holder.dest {
		if i < baseIndex {
			// dest text does not exist in source. so mark this text as remove-diff
			holder.markRemove(value)
			continue
		}
		if i == baseIndex {
			// source text exists in dest. so do nothing.
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	holder.dest = nextHolds
}

func (holder *Holder) markAdd(value Value) {
	holder.diffs.Add(value)
}

func (holder *Holder) markRemove(value Value) {
	holder.diffs.Remove(value)
}

func (holder *Holder) GetDiffs() *Diffs {
	return &holder.diffs
}
