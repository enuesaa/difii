package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffs(t *testing.T) {
	diffs := NewDiffs()
	diffs.Add(*NewValue(2, true, "aaa"))
	diffs.Remove(*NewValue(2, true, "bbb"))

	// TODO: why diffs.ListItems() returns []Diffline ?
	assert.Equal(t, []Diffline{
		*NewDiffline(*NewValue(2, true, "aaa"), Added),
		*NewDiffline(*NewValue(2, true, "bbb"), Removed),
	}, diffs.ListItems())

	expectedHunk := NewHunk()
	expectedHunk.Push(*NewDiffline(*NewValue(2, true, "aaa"), Added))
	expectedHunk.Push(*NewDiffline(*NewValue(2, true, "bbb"), Removed))
	assert.Equal(t, []Hunk{*expectedHunk}, diffs.ListHunks())

	assert.Equal(t, 1, diffs.CountAdd())
	assert.Equal(t, 1, diffs.CountRemove())
}

func TestDiffsWith2Hunks(t *testing.T) {
	diffs := NewDiffs()
	diffs.Add(*NewValue(2, true, "aaa"))
	diffs.Remove(*NewValue(4, true, "bbb"))

	// TODO: why diffs.ListItems() returns []Diffline ?
	assert.Equal(t, []Diffline{
		*NewDiffline(*NewValue(2, true, "aaa"), Added),
		*NewDiffline(*NewValue(4, true, "bbb"), Removed),
	}, diffs.ListItems())

	expectedHunk1 := NewHunk()
	expectedHunk1.Push(*NewDiffline(*NewValue(2, true, "aaa"), Added))
	expectedHunk2 := NewHunk()
	expectedHunk2.Push(*NewDiffline(*NewValue(4, true, "bbb"), Removed))
	assert.Equal(t, []Hunk{*expectedHunk1, *expectedHunk2}, diffs.ListHunks())

	assert.Equal(t, 1, diffs.CountAdd())
	assert.Equal(t, 1, diffs.CountRemove())
}
