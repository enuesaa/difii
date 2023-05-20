package files

import (
	"os"
)

func ListDirs(path string) []string {
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

func IsDirExist(path string) bool {
	// see https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}
