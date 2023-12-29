package cmd

import (
	"fmt"
	"goprojects/taskManager/db"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
	Long:  "Task with the Id given in the argument is marked as compeleted in the TODO list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := convertToInt(args[0])
		err := db.Do(number)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func convertToInt(str string) (int, error) {
	number, err := strconv.Atoi(str)
	if err != nil {
		// Handle error if conversion fails
		fmt.Println("Conversion error:", err)
		return 0, err
	}
	return number, nil
}
