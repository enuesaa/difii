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
	diff := analyzer.Analyze().Render()
	assert.Equal(t, heredoc.Doc(`
	1 - bbbb
	1 + aaaa
	`), diff)
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
	diff := analyzer.Analyze().Render()
	assert.Equal(t, heredoc.Doc(`
	3 - ffffff
	3 + cccccc
	4 + dddddd
	5 + eeeeee
	`), diff)
}

func TestHunkedWithEmptyLine(t *testing.T) {
	// dest の途中に remove-diff があるとき、それ以降の行が diff 扱いされてしまわないよう

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
	diff := analyzer.Analyze().Render()
	assert.Equal(t, heredoc.Doc(`
	3 - 
	`), diff)
}
