package statuschecker

import (
	"fmt"
	"net/http"
)

func CheckUrl(url string, channel chan string) {
	if _, err := http.Get(url); err != nil {
		fmt.Printf("%v might be down!\n", url)
	} else {
		fmt.Printf("%v is up!\n", url)
	}
	channel <- url
}
