package main

import (
	"github.com/enuesaa/difii/pkg/cli"
	"github.com/enuesaa/difii/pkg/repo"
)

func main() {
	prompt := repo.NewPrompt()
	files := repo.NewFiles()

	app := cli.CreateCli(prompt, files)
	app.Execute()
}
