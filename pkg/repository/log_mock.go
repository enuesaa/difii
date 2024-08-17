package repository

import (
	"fmt"
)

type LogMock struct {
	Out string
}

func (repo *LogMock) Printf(format string, a ...any) {
	repo.Out += fmt.Sprintf(format, a...)
}
