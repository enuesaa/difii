package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	renderer := NewMockRenderer()
	input := CliInput {
		CompareDir: "../../testdata/simple-a",
		WorkDir: "../../testdata/simple-b",
		Includes: make([]string, 0),
		Interactive: false,
		Summary: true,
		Inspect: false,
		Apply: false,
	}

	summarySrv := SummaryService{}
	summarySrv.Render(renderer, input)
	assert.Equal(t, "Diffs Summary\n-1 +0 diffs in main.md \n\n", renderer.Out)
}
