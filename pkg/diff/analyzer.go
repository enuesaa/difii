package diff

import (
	"bufio"
	"io"
)

type Analyzer struct {
	source bufio.Scanner
	dest bufio.Scanner
	diffs Diffs
}
func NewAnalyzer(source io.Reader, dest io.Reader) *Analyzer {
	return &Analyzer {
		source: *bufio.NewScanner(source),
		dest: *bufio.NewScanner(dest),
		diffs: *NewDiffs(),
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

func (anly *Analyzer) Analyze() {
	for {
		sourceValue, destValue := anly.next()
		if sourceValue.Has() {
			if destValue.Has() {
				if sourceValue.Text() != destValue.Text() {
					anly.add(sourceValue.Text())
					anly.remove(destValue.Text())
				}
			} else {
				anly.add(sourceValue.Text())
			}
		} else {
			anly.remove(destValue.Text())

			if !destValue.Has() {
				break;
			}
		}
	}
}

func (anly *Analyzer) add(text string) {
	anly.diffs.Add(text)
}

func (anly *Analyzer) remove(text string) {
	anly.diffs.Remove(text)
}

func (anly *Analyzer) GetDiffs() *Diffs {
	return &anly.diffs
}
