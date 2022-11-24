package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	baseDir := "sample"
	var toRename []string
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		return nil
	})
	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		matchResult, _ := match(filename)
		newPath := filepath.Join(dir, matchResult)
		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Println("Error renaming:", err.Error())
		} else {
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
		}
	}
}

// match returns the new file name, or an error if the file name
// did'nt match pattern.
func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didn't match pattern", filename)
	}
	return re.ReplaceAllString(filename, replaceString), nil
}
