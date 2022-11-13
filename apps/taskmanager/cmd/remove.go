package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/romankravchuk/learn-go/apps/taskmanager/storage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove your task.",
	Run:   removeTask,
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

func removeTask(cmd *cobra.Command, args []string) {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse the argument:", arg)
			return
		} else {
			ids = append(ids, id)
		}
	}
	tasks, err := storage.GetTasks()
	if err != nil {
		log.Fatal("Something went wrong:", err)
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		err := storage.DeleteTask(task.Key)
		if err != nil {
			fmt.Printf("Failed to remove task #%d. Error: %s\n", id, err)
		} else {
			fmt.Printf("Removed task #%d\n", id)
		}
	}
}
