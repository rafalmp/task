package cmd

import (
	"fmt"
	"strconv"

	"github.com/rafalmp/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do [flags] <task number(s)>",
	Short: "Mark a task as complete",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskIds []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse arg:", arg)
			} else {
				taskIds = append(taskIds, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("An error occurred: ", err)
			return
		}
		for _, id := range taskIds {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}
	},
}
