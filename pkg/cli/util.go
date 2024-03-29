package cli

import (
	"slices"
	"strings"

	"github.com/enuesaa/difii/pkg/repository"
)

func listTargetFiles(fsio repository.FsioInterface, workDir string, compareDir string) []string {
	list := make([]string, 0)
	list = append(list, fsio.ListFilesRecursively(workDir)...)
	list = append(list, fsio.ListFilesRecursively(compareDir)...)

	list = removeDuplicateFiles(list)
	list = removeGitDir(list)

	return list
}

func removeDuplicateFiles(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func removeGitDir(list []string) []string {
	list = slices.DeleteFunc(list, func(path string) bool {
		return strings.HasPrefix(path, ".git/")
	})
	return list
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
