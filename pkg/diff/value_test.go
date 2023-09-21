package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	value := NewValue(2, true, "aaa")

	assert.Equal(t, 2, value.Line())
	assert.Equal(t, true, value.Has())
	assert.Equal(t, "aaa", value.Text())
}
