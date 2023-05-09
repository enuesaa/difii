package files

import (
	"os"
)

func ListFiles(dir string) []string {
	files, _ := os.ReadDir(dir)
	filenames := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			filenames = append(filenames, file.Name() + "/")
		} else {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
}