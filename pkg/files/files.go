package files

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func ListFilesInDirs(dirs ...string) []string {
	list := make([]string, 0)

	for _, dir := range dirs {
		list = append(list, listFiles(dir)...)
	}
	return removeDuplicateFiles(list)
}

// see https://zenn.dev/orangekame/articles/dad6d0e9382660
func removeDuplicateFiles(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func listFiles(dir string) []string {
	filenames := make([]string, 0)
	filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasPrefix(path, ".git/") {
			return nil
		}
		if file.IsDir() {
			return nil
		} else {
			filenames = append(filenames, path)
		}
		return nil
	})

	return removeRelativePathFromFilenames(filenames, dir+"/")
}

func removeRelativePathFromFilenames(filenames []string, path string) []string {
	converted := make([]string, 0)
	for _, filename := range filenames {
		converted = append(converted, strings.TrimPrefix(filename, path))
	}
	return converted
}

func ReadStream(path string) *os.File {
	f, _ := os.Open(path)

	return f
}
