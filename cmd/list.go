package cmd

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
    list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Open(fmt.Sprintf("%s/.project-setup", home))
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
