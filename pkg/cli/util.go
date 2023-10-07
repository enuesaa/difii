package cli

import (
	"slices"

	"github.com/enuesaa/difii/pkg/repo"
)

func listTargetFiles(fsio repo.FsioInterface, workDir string, compareDir string) []string {
	list := make([]string, 0)
	list = append(list, fsio.ListFiles(workDir)...)
	list = append(list, fsio.ListFiles(compareDir)...)

	return removeDuplicateFiles(list)
}

// see https://zenn.dev/orangekame/articles/dad6d0e9382660
func removeDuplicateFiles(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func filterIncludeFiles(files []string, includes []string) []string {
	list := make([]string, 0)
	for _, filename := range files {
		if slices.Contains(includes, filename) {
			list = append(list, filename)
		}
	}

	return list
}
