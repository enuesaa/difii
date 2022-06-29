package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
	// get `template-name`
	// check is `template-name` in `~/.project-setup`
	// copy dir
	
	rootCmd.AddCommand(installCmd)
}
