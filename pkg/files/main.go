package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(dir string) []string {
	files, _ := os.ReadDir(dir)
	filenames := make([]string, 0)
	for _, file := range files {
		if file.Name() == ".git" {
			continue
		}
		if file.IsDir() {
			filenames = append(filenames, file.Name()+"/")
		} else {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames
}

func ListFilesRecursively(dir string) []string {
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
