package main

import (
	"github.com/enuesaa/difii/pkg/cli/commands"
)

func main() {
	var cli = commands.CreateCli()
	cli.Execute()
}
