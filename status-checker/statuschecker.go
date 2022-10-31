package statuschecker

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	urls := []string{
		"http://go.dev",
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	channel := make(chan string)

	for _, url := range urls {
		go checkUrl(url, channel)
	}

	for url := range channel {
		go func(url string) {
			time.Sleep(2 * time.Second)
			checkUrl(url, channel)
		}(url)
	}
}

func checkUrl(url string, channel chan string) {
	if _, err := http.Get(url); err != nil {
		fmt.Printf("%v might be down!\n", url)
	} else {
		fmt.Printf("%v is up!\n", url)
	}
	channel <- url
}
