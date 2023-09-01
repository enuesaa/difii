package cli

import (
	"fmt"
)

func RecommendInspectFlag(ren RendererInterface, input CliInput) {
	ren.Printf("To inspect diffs:\n")

	if input.IsFileSpecified() {
		onlyflag := ""
		for _, filename := range input.Includes {
			onlyflag = onlyflag + fmt.Sprintf(" --only %s", filename)
		}
		ren.Printf("  difii --workdir %s --compare %s %s --inspect\n", input.BaseDir, input.CompareDir, onlyflag)
	} else {
		ren.Printf("  difii --workdir %s --compare %s --inspect\n", input.BaseDir, input.CompareDir)
	}
	ren.Printf("\n")
}
