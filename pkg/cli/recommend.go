package cli

import (
	"fmt"
)

func RecommendInspectCmd(input CliInput) {
	fmt.Printf("To inspect diffs:\n")
	fmt.Printf("  difii --source %s --dest %s inspect\n", input.SourceDir, input.DestDir)
	fmt.Printf("\n")
}
