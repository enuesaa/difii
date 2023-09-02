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


func TestSummaryForMultiFiles(t *testing.T) {
	cases := []struct{
		workDir string
		compareDir string
		diff string
	} {
		{
			workDir: "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-filename-changed",
			diff: "-0 +8 diffs in changed.md \n-8 +0 diffs in main.md",
		},
	}

    for _, tc := range cases {
		input := CliInput {
			CompareDir: tc.compareDir,
			WorkDir: tc.workDir,
			Includes: make([]string, 0),
			Interactive: false,
			Summary: true,
			Inspect: false,
			Apply: false,
		}
	
		summarySrv := SummaryService{}
		renderer := NewMockRenderer()
		summarySrv.Render(renderer, input)
		assert.Equal(t, "-----------\n\nSummary\n\n" + tc.diff + " \n\n", renderer.Out)
	}
}

