package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/romankravchuk/learn-go/lib/middleware"
)

func main() {

	l := log.New(os.Stdout, "[ak1m1tsu] ", 0)
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)

	router := middleware.New(middleware.NewRecovery())
	router.UseHandler(mux)

	l.Println("listening on http//localhost:3000")
	l.Fatal(http.ListenAndServe(":3000", router))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
}
