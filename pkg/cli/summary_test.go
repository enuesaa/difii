package cli

import (
	"fmt"
	"testing"

	"github.com/enuesaa/difii/pkg/repo"
	"github.com/stretchr/testify/assert"
)

func TestSummaryDiffsCount(t *testing.T) {
	cases := []struct {
		workDir    string
		compareDir string
		diff       string
	}{
		{
			workDir:    "../../testdata/simple-a",
			compareDir: "../../testdata/simple-b",
			diff:       "-0 +1",
		},
		{
			workDir:    "../../testdata/random-a",
			compareDir: "../../testdata/random-b",
			diff:       "-5 +4",
		},
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-b",
			diff:       "-2 +2",
		},
	}

	for _, tc := range cases {
		input := CliInput{
			CompareDir:  tc.compareDir,
			WorkDir:     tc.workDir,
			Includes:    make([]string, 0),
			Interactive: false,
			Task:        TaskSummary,
		}

		summarySrv := SummaryService{}
		fsio := repo.NewFsioMock()
		summarySrv.Render(fsio, input)
		assert.Equal(t, fmt.Sprintf("%s diffs in main.md \n", tc.diff), fsio.Out)
	}
}

func TestSummaryForMultiFiles(t *testing.T) {
	cases := []struct {
		workDir    string
		compareDir string
		diff       string
	}{
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-filename-changed",
			diff: `-0 +8 diffs in changed.md 
-8 +0 diffs in main.md `,
		},
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-sub-files",
			diff: `-0 +0 diffs in main.md 
-0 +2 diffs in sub.md 
-0 +2 diffs in subsub.md `,
		},
	}

	for _, tc := range cases {
		input := CliInput{
			CompareDir:  tc.compareDir,
			WorkDir:     tc.workDir,
			Includes:    make([]string, 0),
			Interactive: false,
			Verbose:     true,
			Task:        TaskSummary,
		}

		summarySrv := SummaryService{}
		fsio := repo.NewFsioMock()
		summarySrv.Render(fsio, input)
		assert.Equal(t, fmt.Sprintf("%s\n", tc.diff), fsio.Out)
	}
}
