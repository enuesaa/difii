package cmd

import (
	"fmt"
	"os"
	"log"
	"io"

	"github.com/spf13/cobra"
	cp "github.com/otiai10/copy"
)

var installCmd = &cobra.Command{
	Use: "install",
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}
var templateName string

func init() {
	installCmd.Flags().StringVar(&templateName, "template-name", "", "template name (required)")
	installCmd.MarkFlagRequired("template-name")
	rootCmd.AddCommand(installCmd)
}

func install() {
	fmt.Printf("template name: %s\n", templateName)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// check `template-name` is in `~/.project-setup`
	dir, err := os.Open(fmt.Sprintf("%s/.project-setup/%s", home, templateName))
	if err != nil {
		fmt.Printf("error! there is no template like: %s\n", templateName)
		os.Exit(0)
	}
	defer dir.Close()

	// check target dir is empty
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("install dir: %s\n", path)

	is_empty := dirIsEmpty(path)
	if (is_empty == false) {
		fmt.Printf("error! install dir is not empty\n")
		os.Exit(0)
	}
	fmt.Printf("install dir is empty")

	cp.Copy(fmt.Sprintf("%s/.project-setup/%s", home, templateName), path)
	os.Exit(0)
}


func dirIsEmpty(name string) (bool) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}
