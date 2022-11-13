package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/romankravchuk/learn-go/apps/taskmanager/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run:   addTask,
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	_, err := storage.CreateTask(task)
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}
	fmt.Printf("Added \"%s\" to your task list.\n", task)
}
