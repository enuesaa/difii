package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffs(t *testing.T) {
	diffs := NewDiffs()
	diffs.MarkAdd(*NewValue(2, true, "aaa"))
	diffs.MarkRemove(*NewValue(2, true, "bbb"))

	// TODO: why diffs.ListItems() returns []Diffline ?
	assert.Equal(t, []Diffline{
		*NewDiffline(2, "aaa", Added),
		*NewDiffline(2, "bbb", Removed),
	}, diffs.ListItems())

	expectedHunk := NewHunk()
	expectedHunk.Push(*NewDiffline(2, "aaa", Added))
	expectedHunk.Push(*NewDiffline(2, "bbb", Removed))
	assert.Equal(t, []Hunk{*expectedHunk}, diffs.ListHunks())

	assert.Equal(t, 1, diffs.CountAdd())
	assert.Equal(t, 1, diffs.CountRemove())
}

func TestDiffsWith2Hunks(t *testing.T) {
	diffs := NewDiffs()
	diffs.MarkAdd(*NewValue(2, true, "aaa"))
	diffs.MarkRemove(*NewValue(4, true, "bbb"))

	// TODO: why diffs.ListItems() returns []Diffline ?
	assert.Equal(t, []Diffline{
		*NewDiffline(2, "aaa", Added),
		*NewDiffline(4, "bbb", Removed),
	}, diffs.ListItems())

	expectedHunk1 := NewHunk()
	expectedHunk1.Push(*NewDiffline(2, "aaa", Added))
	expectedHunk2 := NewHunk()
	expectedHunk2.Push(*NewDiffline(4, "bbb", Removed))
	assert.Equal(t, []Hunk{*expectedHunk1, *expectedHunk2}, diffs.ListHunks())

	assert.Equal(t, 1, diffs.CountAdd())
	assert.Equal(t, 1, diffs.CountRemove())
}
