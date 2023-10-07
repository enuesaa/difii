package main

import (
	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/repo"
)

func main() {
	fsio := repo.NewFsio()
	app := cli.CreateCli(fsio)
	app.Execute()
}
