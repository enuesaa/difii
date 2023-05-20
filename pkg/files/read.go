package files

import (
	"bufio"
	"os"

	"fmt"
)

func ReadStreamWithDiff(sourceDir string, destinationDir string, filename string) {
	source, err := os.Open(sourceDir + "/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()

	destination, err := os.Open(destinationDir + "/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destination.Close()

	sourceScanner := bufio.NewScanner(source)
	destinationScanner := bufio.NewScanner(destination)

	// todo hunked
	// todo when destination holds longer bytes than source
	for sourceScanner.Scan() {
		if destinationScanner.Scan() {
			sourceText := sourceScanner.Text()
			destinationText := destinationScanner.Text()
			if sourceText != destinationText {
				fmt.Printf("+ %s\n", sourceText)
				fmt.Printf("- %s\n", destinationText)
			}
		} else {
			sourceText := sourceScanner.Text()
			fmt.Printf("+ %s\n", sourceText)
		}
	}
}
