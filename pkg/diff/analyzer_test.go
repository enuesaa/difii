package diff

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	source := strings.NewReader("aaaa")
	dest := strings.NewReader("bbbb")

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 1, text: "bbbb", added: false},
		{line: 1, text: "aaaa", added: true},
	}, diff)
}

func TestHunked(t *testing.T) {
	source := strings.NewReader(heredoc.Doc(`
	aaaa
	bbbb
	cccccc
	dddddd
	eeeeee
	gggg
	`))
	dest := strings.NewReader(heredoc.Doc(`
	aaaa
	bbbb
	ffffff
	gggg
	`))

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 3, text: "ffffff", added: false},
		{line: 3, text: "cccccc", added: true},
		{line: 4, text: "dddddd", added: true},
		{line: 5, text: "eeeeee", added: true},
	}, diff)
}

func TestHunkedWithEmptyLine(t *testing.T) {
	source := strings.NewReader(heredoc.Doc(`
	aaaa
	bbbb
	cccccc
	eeeeee
	gggg
	`))
	dest := strings.NewReader(heredoc.Doc(`
	aaaa
	bbbb

	cccccc
	eeeeee
	gggg
	`))

	analyzer := NewAnalyzer(source, dest)
	diff := analyzer.Analyze().ListItems()
	assert.Equal(t, []Diffline{
		{line: 3, text: "", added: false},
	}, diff)
}
