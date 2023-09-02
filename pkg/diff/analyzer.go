package diff

import (
	"bufio"
	"io"
)

type Analyzer struct {
	source bufio.Scanner
	dest   bufio.Scanner
	sourceReading bool
	destReading bool
}

func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer{
		source: *bufio.NewScanner(source),
		dest:   *bufio.NewScanner(dest),
		sourceReading: false,
		destReading: false,
	}
}

// todo refactor
func (anly *Analyzer) next(line int) (Value, Value) {
	sourceHasNext := anly.sourceReading
	destHasNext := anly.destReading

	var sourceNext string
	if anly.source.Scan() {
		sourceNext = anly.source.Text()
		anly.sourceReading = true
		sourceHasNext = true
	} else {
		anly.sourceReading = false
	}

	var destNext string
	if anly.dest.Scan() {
		destNext = anly.dest.Text()
		anly.destReading = true
		destHasNext = true
	} else {
		anly.destReading = false
	}

	return *NewValue(line, sourceHasNext, sourceNext), *NewValue(line, destHasNext, destNext)
}

func (anly *Analyzer) Analyze() *Diffs {
	holder := NewHolder()

	line := 1
	for {
		sourceValue, destValue := anly.next(line)
		if destValue.Has() {
			holder.HoldDest(destValue)
		}
		if sourceValue.Has() {
			holder.HoldSource(sourceValue)
		}
		holder.Flush()

		if !destValue.Has() && !sourceValue.Has() {
			holder.FlushRest()
			break
		}

		line++
	}

	return holder.GetDiffs()
}
