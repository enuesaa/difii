package files

import (
	"os"
)

func ReadStream(path string) *os.File {
	f, _ := os.Open(path)
	// defer f.Close()

	return f
}

// func overwrite
