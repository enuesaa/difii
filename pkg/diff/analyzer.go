package diff

import (
	"bufio"
	"io"
)

type Analyzer struct {
	source        bufio.Scanner
	dest          bufio.Scanner
	sourceReading bool
	destReading   bool
}

func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer{
		source:        *bufio.NewScanner(source),
		dest:          *bufio.NewScanner(dest),
		sourceReading: false,
		destReading:   false,
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
