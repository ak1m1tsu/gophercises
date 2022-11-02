package urlshort

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	YAML string
	JSON string
}

func GetFileData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		Exit(err.Error())
	}
	return string(data)
}

func GetConfig() Config {
	jsonFile := flag.String("json", "", "A json file in the format of \n\t[\n\t\t{\n\t\t\t\"path\":\"/some-url\",\n\t\t\t\"url\":\"https://some-url.com\"\n\t\t},\n\t]")
	yamlFile := flag.String("yaml", "", "A yaml file in the format of \n\t'- path: /some-url'\n\t'  url: https://some-url.com'")
	flag.Parse()
	return Config{
		JSON: *jsonFile,
		YAML: *yamlFile,
	}
}

func DefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func Exit(err string) {
	fmt.Println("ERROR |", err)
	os.Exit(1)
}
