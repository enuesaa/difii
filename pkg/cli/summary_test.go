package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {
	input := CliInput {
		CompareDir: "../../testdata/simple-b",
		WorkDir: "../../testdata/simple-a",
		Includes: make([]string, 0),
		Interactive: false,
		Summary: true,
		Inspect: true,
		Apply: false,
	}

	assert.Equal(t, input.IsCompareDirSelected(), true)
	assert.Equal(t, input.IsWorkDirSelected(), true)
	assert.Equal(t, input.IsFileSpecified(), false)
	assert.Equal(t, input.HasNoOperationFlags(), false)
	assert.Equal(t, input.HasNoGlobalFlags(), false)
	assert.Equal(t, input.Validate(), nil)
}
