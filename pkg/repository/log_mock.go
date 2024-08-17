package repository

import (
	"fmt"
)

//TODO how to access this out field ?
type LogMock struct {
	Out  string
}

func (repo *LogMock) Printf(format string, a ...any) {
	repo.Out += fmt.Sprintf(format, a...)
}
