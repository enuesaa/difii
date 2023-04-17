package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
)

func init() {
	var promenadeCmd = &cobra.Command{
		Use: "promenade",
		Run: func(cmd *cobra.Command, args []string) {
			promenade()
		},
	}

	rootCmd.AddCommand(promenadeCmd)
}

func promenade() {
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
