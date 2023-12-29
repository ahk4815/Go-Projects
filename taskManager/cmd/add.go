package cmd

import (
	"goprojects/taskManager/db"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to our TODO list.",
	Long:  "A new task is added to the inventory of tasks managed by the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Add(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}
	},
}
