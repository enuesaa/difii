package cli

import (
	"fmt"
)

func RecommendInspectCmd(input CliInput) {
	fmt.Printf("To inspect diffs:\n")
	fmt.Printf("  difii --compare %s --workdir %s inspect\n", input.CompareDir, input.WorkDir)
	fmt.Printf("\n")
}
