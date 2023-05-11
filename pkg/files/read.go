package files

import (
	"os"
)

func Read(dir string, filename string) string {
	bytes, err := os.ReadFile(dir + "/" + filename)
	if err != nil {
		return ""
	}
	return string(bytes)
}
