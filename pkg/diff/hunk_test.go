package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunk(t *testing.T) {
	hunk := NewHunk()
	hunk.Push(*NewDiffline(2, "aaa", Added))
	hunk.Push(*NewDiffline(3, "bbb", Removed))

	assert.Equal(t, []Diffline{
		*NewDiffline(2, "aaa", Added),
		*NewDiffline(3, "bbb", Removed),
	}, hunk.ListItems())
}
