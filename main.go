package main

import (
	"github.com/enuesaa/difii/pkg/commands"
)

func main() {
	var cli = commands.CreateCli()
	cli.Execute()
}
