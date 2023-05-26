package diff

import (
	"bufio"
	"io"
)

type Analyzer struct {
	source bufio.Scanner
	dest bufio.Scanner
}
func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer {
		source: *bufio.NewScanner(source),
		dest: *bufio.NewScanner(dest),
	}
}

func (anly *Analyzer) next() (*Value, *Value) {
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

	return &Value{ has: sourceHasNext, text: sourceNext }, &Value{ has: destHasNext, text: destNext }
}

func (anly *Analyzer) Analyze() *Diffs {
	holder := NewHolder()

	for {
		sourceValue, destValue := anly.next()
		if destValue.Has() {
			holder.Hold(destValue.Text())
		}

		if sourceValue.Has() {
			holder.HoldAdd(sourceValue.Text())
		}

		holder.Flush()

		if !destValue.Has() && !sourceValue.Has() {
			holder.FlushRest()
			break;
		}
	}

	return holder.GetDiffs()
}
