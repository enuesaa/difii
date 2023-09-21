package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffline(t *testing.T) {
	diffline := NewDiffline(*NewValue(2, true, "aaa"), Added)

	assert.Equal(t, true, diffline.Added())
	assert.Equal(t, 2, diffline.Line())
	assert.Equal(t, "aaa", diffline.Text())
}
