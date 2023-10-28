package cli

import (
	"testing"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	fsio := repo.NewFsioMock()

	input := CliInput{
		CompareDir:  "../../testdata/simple-b",
		WorkDir:     "../../testdata/simple-a",
		Includes:    make([]string, 0),
		Interactive: false,
		Inspect:     true,
		Apply:       false,
	}

	assert.Equal(t, input.IsCompareDirSelected(), true)
	assert.Equal(t, input.IsWorkDirSelected(), true)
	assert.Equal(t, input.IsFileSpecified(), false)
	assert.Equal(t, input.HasNoOperationFlags(), false)
	assert.Equal(t, input.HasNoGlobalFlags(), false)
	assert.Equal(t, input.Validate(fsio), nil)
}
