package cmd

import (
	"fmt"
	"strconv"

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
		fmt.Printf("Marking task(s) as complete: %v\n", taskIds)
	},
}
