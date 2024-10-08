package cli

import (
	"fmt"
	"testing"

	"github.com/enuesaa/difii/pkg/repository"
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
			diff: `
main.md:2   + b`,
		},
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-sub-files",
			diff: `
sub.md:1   + sub file
sub.md:2   + 
subsub.md:1   + subsub
subsub.md:2   + `,
		},
		{
			workDir:    "../../testdata/tourism-a",
			compareDir: "../../testdata/tourism-nested-files",
			diff: `
nested/main.md:1   + this is nested file.
nested/main.md:2   + `,
		},
	}

	for _, tc := range cases {
		input := CliInput{
			CompareDir:  tc.compareDir,
			WorkDir:     tc.workDir,
			Includes:    make([]string, 0),
			Interactive: true,
			Task:        TaskInspect,
		}

		repos := repository.NewMock()
		inspectSrv := NewInspectService(repos)
		inspectSrv.Render(input)
		assert.Equal(t, fmt.Sprintf("----- Inspect -----%s\n", tc.diff), repos.Log.(*repository.LogMock).Out)
	}
}
