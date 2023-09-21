package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffline(t *testing.T) {
	diffline := NewDiffline(*NewValue(1, true, "aaa"), Added)

	assert.Equal(t, diffline.Added(), true)
	assert.Equal(t, diffline.Line(), 1)
	assert.Equal(t, diffline.Text(), "aaa")
}
