package main

import (
	"github.com/enuesaa/difii/pkg/cli"
)

func main() {
	app := cli.CreateCli()
	app.Execute()
}
