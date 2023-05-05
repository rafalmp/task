package cmd

import (
	"fmt"
	"strings"

	"github.com/rafalmp/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [flags] <task description>",
	Short: "Add a new task to your list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("An error occurred: ", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}
