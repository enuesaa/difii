package files

import (
	"os"
	"bufio"

	"fmt"
)

func Read(dir string, filename string) string {
	bytes, err := os.ReadFile(dir + "/" + filename)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ReadStream(dir string, filename string) {
	// see https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	// see https://zenn.dev/hsaki/books/golang-io-package/viewer/bufio
	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


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
