package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "difii",
	Run: rootCmdHandler,
}

func rootCmdHandler(cmd *cobra.Command, args []string)  {
	fmt.Println("aa")
}
