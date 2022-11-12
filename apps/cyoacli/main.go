package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/romankravchuk/learn-go/lib/cyoa"
)

var (
	filename        string
	ErrFileNotFound error
)

func init() {
	flag.StringVar(&filename, "file", "", "The JSON file with the CYOA story.")
	ErrFileNotFound = errors.New("can not find the file")
}

func main() {
	// TODO: Write cli application
	flag.Parse()
	if filename == "" {
		exit(ErrFileNotFound)
	}

	f, err := os.Open(filename)
	exit(err)

	story, err := cyoa.JsonStory(f)
	exit(err)

	printParagraphs(story["intro"])
}

func printParagraphs(chapter cyoa.Chapter) {
	fmt.Printf("%s\n\n", chapter.Title)
	for _, paragraph := range chapter.Paragraphs {
		fmt.Printf("%s\n", paragraph)
	}
	fmt.Printf("\nOptions:\n")
	for i, option := range chapter.Options {
		fmt.Println(i+1, " - ", option.Text)
	}
}

func exit(err error) {
	if err != nil {
		log.Fatalf("ERROR | %v", err)
	}
}
