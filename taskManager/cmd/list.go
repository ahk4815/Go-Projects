package cmd

import (
	"goprojects/taskManager/db"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays the tasks in out TODO list.",
	Long:  "The incomplete tasks in our TODO list is displayed by the command.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.List()
		if err != nil {
			log.Fatal(err)
		}
	},
}
