package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(_ *cobra.Command, _ []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete, Take a break!")
			return
		}
		fmt.Println("Tasks:")
		for i, task := range tasks {
			var tick string = " "
			if task.Value.Complete {
				tick = "X"
			}
			fmt.Printf("%d: [%s] %s\n", i+1, tick, task.Value.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
