package diff

import (
	"bufio"
	"io"
	"golang.org/x/exp/slices"
)

type Analyzer struct {
	source        bufio.Scanner
	dest          bufio.Scanner
	sourceReading bool
	destReading   bool
	sourceValues  []Value // tmp store
	destValues    []Value // tmp store
	diffs         Diffs
}

func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer{
		source:        *bufio.NewScanner(source),
		dest:          *bufio.NewScanner(dest),
		sourceReading: false,
		destReading:   false,
		sourceValues:  make([]Value, 0),
		destValues:    make([]Value, 0),
		diffs:         *NewDiffs(),
	}
}

// todo refactor
func (anly *Analyzer) next(line int) (Value, Value) {
	// bufio.Scanner misses last empty line.
	// to prevent this, use anly.sourceReading as reading status.
	sourceValue := *NewValue(line, anly.sourceReading, "")
	destValue := *NewValue(line, anly.destReading, "")

	anly.sourceReading = anly.source.Scan()
	if anly.sourceReading {
		sourceValue = *NewValue(line, true, anly.source.Text())
	}

	anly.destReading = anly.dest.Scan()
	if anly.destReading {
		destValue = *NewValue(line, true, anly.dest.Text())
	}

	return sourceValue, destValue
}

func (anly *Analyzer) Analyze() *Diffs {
	line := 1
	for {
		sourceValue, destValue := anly.next(line)
		if destValue.Has() {
			anly.holdDest(destValue)
		}
		if sourceValue.Has() {
			anly.holdSource(sourceValue)
		}
		anly.flush()

		if !destValue.Has() && !sourceValue.Has() {
			anly.flushRest()
			break
		}

		line++
	}

	return &anly.diffs
}

func (anly *Analyzer) holdDest(value Value) {
	anly.destValues = append(anly.destValues, value)
}

func (anly *Analyzer) holdSource(value Value) {
	anly.sourceValues = append(anly.sourceValues, value)
}

func (anly *Analyzer) getHoldDestIndex(source Value) int {
	// see https://stackoverflow.com/questions/38654383/how-to-search-for-an-element-in-a-golang-slice
	i := slices.IndexFunc(anly.destValues, func(value Value) bool {
		return value.Text() == source.Text()
	})
	return i
}

func (anly *Analyzer) flush() {
	for i, sourceValue := range anly.sourceValues {
		matched := anly.getHoldDestIndex(sourceValue)
		if matched == -1 {
			continue
		}

		anly.rebaseDest(matched)
		anly.rebaseSource(i)
	}
}

func (anly *Analyzer) flushRest() {
	for _, value := range anly.destValues {
		anly.markRemove(value)
	}
	for _, value := range anly.sourceValues {
		anly.markAdd(value)
	}
}

func (anly *Analyzer) rebaseSource(baseIndex int) {
	nextHolds := make([]Value, 0)
	for i, value := range anly.sourceValues {
		if i < baseIndex {
			// source text does not exist in dest. so mark this text as add-diff
			anly.markAdd(value)
			continue
		}
		if i == baseIndex {
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	anly.sourceValues = nextHolds
}

func (anly *Analyzer) rebaseDest(baseIndex int) {
	nextHolds := make([]Value, 0)
	for i, value := range anly.destValues {
		if i < baseIndex {
			// dest text does not exist in source. so mark this text as remove-diff
			anly.markRemove(value)
			continue
		}
		if i == baseIndex {
			// source text exists in dest. so do nothing.
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	anly.destValues = nextHolds
}

func (anly *Analyzer) markAdd(value Value) {
	anly.diffs.Add(value)
}

func (anly *Analyzer) markRemove(value Value) {
	anly.diffs.Remove(value)
}
