package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
)

func init() {
	var promenadeCmd = &cobra.Command{
		Use: "promenade",
		Run: promenadeCmdHandler,
	}

	RootCmd.AddCommand(promenadeCmd)
}

func promenadeCmdHandler(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "selector",
		Items: []string{
			"a",
			"b",
			"c",
		},
	}
	
	i, result, err := prompt.Run()
	fmt.Printf("%+v", i)
	fmt.Printf("%s", result)
	fmt.Printf("%s", err)
}
