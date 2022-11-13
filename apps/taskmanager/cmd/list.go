package cmd

import (
	"fmt"
	"log"

	"github.com/romankravchuk/learn-go/apps/taskmanager/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run:   getTasks,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func getTasks(cmd *cobra.Command, args []string) {
	tasks, err := storage.GetTasks()
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("You have no tasks to complete! Why not to take vacation? 😸")
		return
	}
	fmt.Println("You have following tasks:")
	for i, task := range tasks {
		fmt.Printf("%d - %s\n", i+1, task.Value)
	}
}
