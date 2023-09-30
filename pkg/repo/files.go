package repo

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type FilesInterface interface {
	ListDirs(path string) []string
	IsDirOrFileExist(path string) bool
	ListFilesInDirs(dirs ...string) []string
	ReadStream(path string) *os.File
	FilterFiles(files []string, includes []string) []string
}

type Files struct {}

func NewFiles() *Files {
	return &Files{}
}

func (repo *Files) ListDirs(path string) []string {
	dirs := make([]string, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		return dirs
	}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}

func (repo *Files) IsDirOrFileExist(path string) bool {
	// see https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (repo *Files) ListFilesInDirs(dirs ...string) []string {
	list := make([]string, 0)

	for _, dir := range dirs {
		list = append(list, repo.listFiles(dir)...)
	}
	return repo.removeDuplicateFiles(list)
}

// see https://zenn.dev/orangekame/articles/dad6d0e9382660
func (repo *Files) removeDuplicateFiles(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func (repo *Files) listFiles(dir string) []string {
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

	return repo.removeRelativePathFromFilenames(filenames, dir+"/")
}

func (repo *Files) removeRelativePathFromFilenames(filenames []string, path string) []string {
	converted := make([]string, 0)
	for _, filename := range filenames {
		converted = append(converted, strings.TrimPrefix(filename, path))
	}
	return converted
}

func (repo *Files) ReadStream(path string) *os.File {
	f, _ := os.Open(path)

	return f
}

func (repo *Files) FilterFiles(files []string, includes []string) []string {
	list := make([]string, 0)
	for _, filename := range files {
		if slices.Contains(includes, filename) {
			list = append(list, filename)
		}
	}

	return list
}
