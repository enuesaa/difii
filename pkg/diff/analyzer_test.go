package diff

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	source := strings.NewReader("aaaa")
	dest := strings.NewReader("bbbb")

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 1, text: "bbbb", differ: Removed},
		{line: 1, text: "aaaa", differ: Added},
	}, diff)
}

func TestHunked(t *testing.T) {
	source := strings.NewReader(`
aaaa
bbbb
cccccc
dddddd
eeeeee
gggg
`)
	dest := strings.NewReader(`
aaaa
bbbb
ffffff
gggg
`)

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 3, text: "ffffff", differ: Removed},
		{line: 3, text: "cccccc", differ: Added},
		{line: 4, text: "dddddd", differ: Added},
		{line: 5, text: "eeeeee", differ: Added},
	}, diff)
}

func TestHunkedWithEmptyLine(t *testing.T) {
	source := strings.NewReader(`
aaaa
bbbb
cccccc
eeeeee
gggg
`)
	dest := strings.NewReader(`
aaaa
bbbb

cccccc
eeeeee
gggg
`)

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 3, text: "", differ: Removed},
	}, diff)
}
