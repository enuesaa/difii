package diff

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Diff(source string, destination string) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(source, destination, false)
	for _, diff := range diffs {
		if diff.Type == diffmatchpatch.DiffEqual {
			continue
		}
		arr := make([]diffmatchpatch.Diff, 1)
		arr = append(arr, diff)
		fmt.Println(dmp.DiffPrettyText(arr))
	}
}
