package urlshort

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func Run() {
	config := getConfig()

	mux := defaultMux()

	pathsToUrls := map[string]string{}
	mapHandler := MapHandler(pathsToUrls, mux)

	yaml := getFileData(config.YAML)
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		exit(err.Error())
	}

	json := getFileData(config.JSON)
	jsonHandler, err := JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Starting the server on localhost:8080")
	http.ListenAndServe(":8080", jsonHandler)
}

type Config struct {
	YAML string
	JSON string
}

func getFileData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		exit(err.Error())
	}
	return string(data)
}

func getConfig() Config {
	jsonFile := flag.String("json", "", "A json file in the format of \n\t[\n\t\t{\n\t\t\t\"path\":\"/some-url\",\n\t\t\t\"url\":\"https://some-url.com\"\n\t\t},\n\t]")
	yamlFile := flag.String("yaml", "", "A yaml file in the format of \n\t'- path: /some-url'\n\t'  url: https://some-url.com'")
	flag.Parse()
	return Config{
		JSON: *jsonFile,
		YAML: *yamlFile,
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func exit(err string) {
	fmt.Println("ERROR |", err)
	os.Exit(1)
}
