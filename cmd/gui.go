package cmd

import (
	"fmt"
	"log"
	"os"
	"task/db"

	"github.com/marcusolsson/tui-go"
	"github.com/spf13/cobra"
)

func fillTaskTable(taskList *tui.Table, tasks []db.Task) {
	for i, task := range tasks {
		var tick string = " "
		if task.Value.Complete {
			tick = "X"
		}
		taskList.AppendRow(
			tui.NewLabel(fmt.Sprint(i+1)),
			tui.NewLabel(fmt.Sprintf("[%s]", tick)),
			tui.NewLabel(task.Value.Text),
		)
	}
}

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Launch GUI",
	Run: func(_ *cobra.Command, _ []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Failed to get tasks")
			log.Fatal(err)
		}

		sidebar := tui.NewVBox(
			tui.NewLabel("All Tasks"),
			tui.NewLabel("Uncompleted Tasks"),
			tui.NewLabel("Completed Tasks"),
			tui.NewSpacer(),
		)
		sidebar.SetBorder(true)

		taskList := tui.NewTable(2, 0)
		fillTaskTable(taskList, tasks)

		taskScroll := tui.NewScrollArea(taskList)
		taskBox := tui.NewVBox(taskScroll)
		taskBox.SetBorder(true)

		input := tui.NewEntry()
		input.SetFocused(true)
		input.SetSizePolicy(tui.Expanding, tui.Maximum)

		inputBox := tui.NewHBox(input)
		inputBox.SetBorder(true)
		inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

		taskSection := tui.NewVBox(taskBox, inputBox)
		taskSection.SetSizePolicy(tui.Expanding, tui.Expanding)

		input.OnSubmit(func(entry *tui.Entry) {
			task := entry.Text()
			_, err := db.CreateTask(task)
			if err != nil {
				log.Fatal(err)
			}
			tasks, err := db.AllTasks()
			if err != nil {
				log.Fatal(err)
			}
			taskList.RemoveRows()
			fillTaskTable(taskList, tasks)
		})

		root := tui.NewHBox(sidebar, taskSection)
		ui, err := tui.New(root)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		ui.SetKeybinding("Esc", func() { ui.Quit() })

		log.Fatal(ui.Run())
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
