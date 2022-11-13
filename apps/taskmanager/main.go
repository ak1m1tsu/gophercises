package main

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/romankravchuk/learn-go/apps/taskmanager/cmd"
	"github.com/romankravchuk/learn-go/apps/taskmanager/storage"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, ".tasks.db")
	must(storage.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
