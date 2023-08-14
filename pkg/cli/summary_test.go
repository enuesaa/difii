package cli

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	t.Helper()
	realStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	input := CliInput {
		CompareDir: "../../testdata/aa-simple",
		WorkDir: "../../testdata/aa-simpler",
		Includes: make([]string, 0),
		Interactive: false,
		Summary: true,
		Inspect: false,
		Apply: false,
	}
	ShowDiffsSummary(input)

	w.Close()
	os.Stdout = realStdout

	buf, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("failed")
		return
	}
	assert.Equal(t, "Diffs Summary\n          -1           +0 diffs in aa.txt \n\n", string(buf))
}
