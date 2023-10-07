package cli

import (
	"slices"

	"github.com/enuesaa/difii/pkg/repo"
)

func listTargetFiles(files repo.FilesInterface, workDir string, compareDir string) []string {
	list := make([]string, 0)
	list = append(list, files.ListFiles(workDir)...)
	list = append(list, files.ListFiles(compareDir)...)

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
