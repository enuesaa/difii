package repo

import (
	"fmt"
)

type FsioMock struct {
	Fsio // test cases depend on real filesystem now, so extends this struct.
	Out  string
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
