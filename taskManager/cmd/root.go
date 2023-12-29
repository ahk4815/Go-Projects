package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Basic Task CLI",
	Long:  "task is a CLI application to manage your TODO tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please choose a command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
