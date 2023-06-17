package main

import (
	"github.com/enuesaa/difii/pkg/cli/command"
)

func main() {
	var cli = command.CreateCli()
	cli.Execute()
}
