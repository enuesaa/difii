package diff

import (
	"bufio"
	"io"
)

type Analyzer struct {
	source bufio.Scanner
	dest   bufio.Scanner
}

func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer{
		source: *bufio.NewScanner(source),
		dest:   *bufio.NewScanner(dest),
	}
}

func (anly *Analyzer) next(line int) (Value, Value) {
	sourceHasNext := anly.source.Scan()
	destHasNext := anly.dest.Scan()

	var sourceNext string
	if sourceHasNext {
		sourceNext = anly.source.Text()
	}
	var destNext string
	if destHasNext {
		destNext = anly.dest.Text()
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
