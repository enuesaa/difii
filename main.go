package main

import (
	"log"

	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/repo"
)

func init() {
	log.SetFlags(0)
}

func main() {
	fsio := repo.NewFsio()
	app := cli.CreateCli(fsio)
	app.Execute()
}
