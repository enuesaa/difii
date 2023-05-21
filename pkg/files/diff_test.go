package files

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/stretchr/testify/assert"
)

func TestCheckin(t *testing.T) {
	source := strings.NewReader("aaaa")
	dest := strings.NewReader("bbbb")

	diff := Diff(source, dest)
	assert.Equal(t, heredoc.Doc(`
	+ aaaa
	- bbbb
	`), diff.String())
}