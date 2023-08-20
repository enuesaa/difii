package cli

import (
	"fmt"
	"strings"
)

func ShowInfo(input CliInput) {
	fmt.Printf("---\n")
	fmt.Printf("workdir    ... %s\n", input.WorkDir)
	fmt.Printf("comparedir ... %s\n", input.CompareDir)

	if len(input.Includes) > 0 {
		fmt.Printf("only \n - %s\n", strings.Join(input.Includes, "\n - "))
	}
	fmt.Printf("---\n")
}
