package main

import (
	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/repository"
)

func main() {
	repos := repository.New()
	app := cli.CreateCli(repos)
	app.Execute()
}
