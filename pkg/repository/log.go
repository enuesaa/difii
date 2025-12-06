package repository

import (
	"log"
)

func init() {
	log.SetFlags(0)
}

type LogInterface interface {
	Printf(format string, a ...any)
}

type Log struct{}

func (repo *Log) Printf(format string, a ...any) {
	log.Printf(format, a...)
}
