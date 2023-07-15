package cli

import (
	"fmt"
)

func RecommendInspectCmd(input CliInput) {
	fmt.Printf("To inspect diffs:\n")

	if input.IsFileSpecified() {
		onlyflag := ""
		for _, filename := range input.Includes {
			onlyflag = onlyflag + fmt.Sprintf(" --only %s", filename)
		}
		fmt.Printf("  difii --workdir %s --compare %s %s inspect\n", input.WorkDir, input.CompareDir, onlyflag)
	} else {
		fmt.Printf("  difii --workdir %s --compare %s inspect\n", input.WorkDir, input.CompareDir)
	}
	fmt.Printf("\n")
}
