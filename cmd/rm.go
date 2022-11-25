package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task from list",
	Run: func(_ *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument:", arg)
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark task \"%d\" as completed, Error; %s\n", id, err)
			} else {
				fmt.Printf("Marked task \"%d\" as completed\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
