package main

import (
	"log"

	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/repository"
)

func init() {
	log.SetFlags(0)
}

func main() {
	repos := repository.New()
	app := cli.CreateCli(repos)
	app.Execute()
}
