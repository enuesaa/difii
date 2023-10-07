package repo

import (
	"fmt"
	"os"
)

type FsioMock struct {
	Out string
}

func NewFsioMock() *FsioMock {
	return &FsioMock{}
}
func (fsio *FsioMock) Printf(format string, a ...any) {
	fsio.Out += fmt.Sprintf(format, a...)
}
func (fsio *FsioMock) Confirm(message string) bool {
	return true
}
func (fsio *FsioMock) SelectCompareDir() string {
	return "./"
}
func (fsio *FsioMock) IsDirOrFileExist(path string) bool {
	return true
}
func (fsio *FsioMock) ListFiles(dir string) []string {
	return make([]string, 0)
}
func (fsio *FsioMock) ReadStream(path string) *os.File {
	return nil
}
