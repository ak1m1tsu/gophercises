package main

import (
	"time"

	statuschecker "github.com/romankravchuk/learn-go/lib/status-checker"
)

func main() {
	urls := []string{
		"http://go.dev",
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	channel := make(chan string)

	for _, url := range urls {
		go statuschecker.CheckUrl(url, channel)
	}

	for url := range channel {
		go func(url string) {
			time.Sleep(2 * time.Second)
			statuschecker.CheckUrl(url, channel)
		}(url)
	}
}
