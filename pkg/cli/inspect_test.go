package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	cases := []struct {
		workDir    string
		compareDir string
		diff       string
	}{
		{
			workDir:    "../../testdata/simple-a",
			compareDir: "../../testdata/simple-b",
			diff:       "main.md:2	+ b",
		},
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-sub-files",
			diff:       "sub.md:1\t+ sub file\nsub.md:2\t+ \n\nsubsub.md:1\t+ subsub\nsubsub.md:2\t+ ",
		},
	}

	for _, tc := range cases {
		input := CliInput{
			CompareDir:  tc.compareDir,
			WorkDir:     tc.workDir,
			Includes:    make([]string, 0),
			Interactive: false,
			Summary:     true,
			Inspect:     false,
			Apply:       false,
		}

		inspectSrv := InspectService{}
		renderer := NewMockRenderer()
		inspectSrv.Render(renderer, input)
		assert.Equal(t, "-----------\n\nInspect\n\n"+tc.diff+"\n\n", renderer.Out)
	}
}
