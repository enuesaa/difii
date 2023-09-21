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
	diffs         Diffs
}

func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer{
		source:        *bufio.NewScanner(source),
		dest:          *bufio.NewScanner(dest),
		sourceReading: false,
		destReading:   false,
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
	destVals   := make([]Value, 0)
	sourceVals := make([]Value, 0)

	line := 1
	for {
		sourceVal, destVal := anly.next(line)
		if destVal.Has() {
			destVals = append(destVals, destVal)
		}
		if sourceVal.Has() {
			sourceVals = append(sourceVals, sourceVal)
		}
		destVals, sourceVals = anly.flush(destVals, sourceVals)

		if !destVal.Has() && !sourceVal.Has() {
			anly.flushRest(destVals, sourceVals)
			break
		}

		line++
	}

	return &anly.diffs
}

func (anly *Analyzer) flush(destVals []Value, sourceVals []Value) ([]Value, []Value) {
	for sourceIndex, sourceValue := range sourceVals {
		destIndex := slices.IndexFunc(destVals, func(value Value) bool {
			return value.Text() == sourceValue.Text()
		})
		if destIndex == -1 {
			continue
		}

		destVals = anly.rebaseDestVals(destVals, destIndex)
		sourceVals = anly.rebaseSourceVals(sourceVals, sourceIndex)
	}
	return destVals, sourceVals
}

func (anly *Analyzer) flushRest(destVals []Value, sourceVals []Value) {
	for _, value := range destVals {
		anly.diffs.MarkRemove(value)
	}
	for _, value := range sourceVals {
		anly.diffs.MarkAdd(value)
	}
}

func (anly *Analyzer) rebaseDestVals(destVals []Value, baseIndex int) []Value {
	nextHolds := make([]Value, 0)
	for i, value := range destVals {
		if i < baseIndex {
			// dest text does not exist in source. so mark this text as remove-diff
			anly.diffs.MarkRemove(value)
			continue
		}
		if i == baseIndex {
			// source text exists in dest. so do nothing.
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	return nextHolds
}

func (anly *Analyzer) rebaseSourceVals(sourceVals []Value, baseIndex int) []Value {
	nextHolds := make([]Value, 0)
	for i, value := range sourceVals {
		if i < baseIndex {
			// source text does not exist in dest. so mark this text as add-diff
			anly.diffs.MarkAdd(value)
			continue
		}
		if i == baseIndex {
			continue
		}
		nextHolds = append(nextHolds, value)
	}
	return nextHolds
}
