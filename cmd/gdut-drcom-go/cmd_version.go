package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("gdut-drcom-go %s (build from %s)", Version, Author))
	},
}

func init() {
	RootCommand.AddCommand(versionCommand)
}
