package repository

import (
	"fmt"
)

type LogInterface interface {
	Printf(format string, a ...any)
}

type Log struct{}

func (repo *Log) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}
