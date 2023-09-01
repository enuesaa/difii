package cli


import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {
	cases := []struct{
		baseDir string
		compareDir string
		diff string
	} {
		{
			baseDir: "../../testdata/simple-a",
			compareDir: "../../testdata/simple-b",
			diff: "-0 +1",
		},
		{
			baseDir: "../../testdata/random-a",
			compareDir: "../../testdata/random-b",
			diff: "-5 +4",
		},
		{
			baseDir: "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-b",
			diff: "-2 +2",
		},
	}

    for _, tc := range cases {
		input := CliInput {
			CompareDir: tc.compareDir,
			BaseDir: tc.baseDir,
			Includes: make([]string, 0),
			Interactive: false,
			Summary: true,
			Inspect: false,
			Apply: false,
		}
	
		summarySrv := SummaryService{}
		renderer := NewMockRenderer()
		summarySrv.Render(renderer, input)
		assert.Equal(t, "-----------\n\nSummary\n\n" + tc.diff + " diffs in main.md \n\n", renderer.Out)
	}
}
