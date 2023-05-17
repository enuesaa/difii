package cli

import (
	"fmt"

	"github.com/enuesaa/difii/pkg/files"
)

func DiffFiles(sourceDir string, destinationDir string) {
	fmt.Printf("source dir: %s \n", sourceDir)
	fmt.Printf("destination dir: %s \n", destinationDir)
	fmt.Println("")

	sourcefiles := files.ListFilesRecursively(sourceDir)
	files.ReadStreamWithDiff(sourceDir, destinationDir, sourcefiles[0])
}
