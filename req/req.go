package req

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))

	return len(p), nil
}

func PrintBodyData(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR |", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
