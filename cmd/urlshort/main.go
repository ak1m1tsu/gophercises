package main

import (
	"fmt"
	"net/http"

	"github.com/romankravchuk/learn-go/urlshort"
)

func main() {
	config := urlshort.GetConfig()

	mux := urlshort.DefaultMux()

	pathsToUrls := map[string]string{}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := urlshort.GetFileData(config.YAML)
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		urlshort.Exit(err.Error())
	}

	json := urlshort.GetFileData(config.JSON)
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		urlshort.Exit(err.Error())
	}

	fmt.Println("Starting the server on localhost:8080")
	http.ListenAndServe(":8080", jsonHandler)
}
