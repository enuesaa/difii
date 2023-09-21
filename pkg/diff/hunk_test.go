package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunk(t *testing.T) {
	hunk := NewHunk()
	hunk.Push(*NewDiffline(*NewValue(2, true, "aaa"), Added))
	hunk.Push(*NewDiffline(*NewValue(3, true, "bbb"), Removed))

	assert.Equal(t, []Diffline{
		*NewDiffline(*NewValue(2, true, "aaa"), Added),
		*NewDiffline(*NewValue(3, true, "bbb"), Removed),
	}, hunk.ListItems())
}
