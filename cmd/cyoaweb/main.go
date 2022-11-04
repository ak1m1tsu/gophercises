package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/romankravchuk/learn-go/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web application on.")
	filename := flag.String("file", "chapters.json", "The JSON file with the CYOA story.")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		exit(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		exit(err)
	}

	ch := cyoa.NewChapterHandler(story)
	fmt.Printf("Starting the web application on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), ch))
}

func exit(err error) {
	fmt.Printf("ERROR | %v", err.Error())
	os.Exit(1)
}
