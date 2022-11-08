package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/romankravchuk/learn-go/link"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "The HTML file")
}

func main() {
	flag.Parse()
	data, err := os.ReadFile(filename)
	if err != nil {
		exit(err)
	}

	r := strings.NewReader(string(data))

	links, err := link.Parse(r)
	if err != nil {
		exit(err)
	}

	fmt.Printf("%+v\n", links)
}

func exit(err error) {
	log.Printf("%v", err)
	os.Exit(1)
}
