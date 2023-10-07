package repo

import (
	"os"
	"path/filepath"
	"strings"
)

type FilesInterface interface {
	IsDirOrFileExist(path string) bool
	ListFiles(dir string) []string
	ReadStream(path string) *os.File
}

type Files struct{}

func NewFiles() *Files {
	return &Files{}
}

func (repo *Files) IsDirOrFileExist(path string) bool {
	// see https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (repo *Files) ListFiles(dir string) []string {
	filenames := make([]string, 0)
	filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO: move 
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
