package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Diff(source io.Reader, dest io.Reader) bytes.Buffer {
	var ret bytes.Buffer

	sourceScanner := bufio.NewScanner(source)
	destScanner := bufio.NewScanner(dest)

	// todo hunked
	// todo when destination holds longer bytes than source
	for sourceScanner.Scan() {
		if destScanner.Scan() {
			sourceText := sourceScanner.Text()
			destText := destScanner.Text()
			if sourceText != destText {
				ret.WriteString(fmt.Sprintf("+ %s\n", sourceText))
				ret.WriteString(fmt.Sprintf("- %s\n", destText))
			}
		} else {
			sourceText := sourceScanner.Text()
			ret.WriteString(fmt.Sprintf("+ %s\n", sourceText))
		}
	}

	return ret
}
