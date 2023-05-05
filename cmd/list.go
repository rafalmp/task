package cmd

import (
	"fmt"
	"os"

	"github.com/rafalmp/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("An error occurred: ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("Nothing to do, go get some rest!")
			return
		}
		fmt.Println("Your task list:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}
