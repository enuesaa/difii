package files

import (
	"golang.org/x/exp/slices"
)

func FilterFiles(files []string, includes []string) []string {
	list := make([]string, 0)
	for _, filename := range files {
		if slices.Contains(includes, filename) {
			list = append(list, filename)
		}
	}

	return list
}
