package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

var dbPath string

func init() {
	home, _ := homedir.Dir()
	dbPath = filepath.Join(home, "test.db")
}

func TestInit(t *testing.T) {
	err := Init(dbPath)
	if err != nil {
		t.Fatalf("Expected DB file on \"%s\", but got error: %s", dbPath, err)
	}
	os.Remove(dbPath)
}

func TestCreateTask(t *testing.T) {
	Init(dbPath)
	id, err := CreateTask("test task")
	if id == -1 {
		t.Fatalf("Expected to create task, but got error: %s", err)
	}
	os.Remove(dbPath)
}

func TestGetTasks(t *testing.T) {
	Init(dbPath)
	CreateTask("test task")
	tasks, _ := GetTasks()
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, but got %d", len(tasks))
	}
	os.Remove(dbPath)
}

func TestDeleteTasks(t *testing.T) {
	Init(dbPath)
	CreateTask("task")
	err := DeleteTask(1)
	if err != nil {
		t.Fatalf("Expected to delete task, but got error: %s", err)
	}
	os.Remove(dbPath)
}
