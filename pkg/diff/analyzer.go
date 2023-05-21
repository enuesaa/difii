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

func (anly *Analyzer) sourceHasNext() bool {
	return anly.source.Scan()
}
func (anly *Analyzer) destHasNext() bool {
	return anly.dest.Scan()
}
func (anly *Analyzer) sourceText() string {
	return anly.source.Text()
}
func (anly *Analyzer) destText() string {
	return anly.dest.Text()
}

func (anly *Analyzer) Analyze() {
	for anly.sourceHasNext() {
		if anly.destHasNext() {
			if anly.sourceText() != anly.destText() {
				anly.diffs.Add(anly.sourceText())
				anly.diffs.Remove(anly.destText())
			}
		} else {
			anly.diffs.Add(anly.sourceText())
		}
	}
}

func (anly *Analyzer) GetDiffs() *Diffs {
	return &anly.diffs
}
