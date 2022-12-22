package main

import (
	"fmt"
	"io"
	"net/http"
)

// webWriter struct
type webWriter struct{}

// Write method
func (webWriter) Write(p []byte) (int, error) {
	// printing p(bytes) as string
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	// taking res from google.com
	res, err := http.Get("https://www.google.com")

	// validating errors
	if err != nil {
		fmt.Println(err)
	}

	// creating a new webWriter
	w := webWriter{}

	// copying res.Body (reader) to w(writer)
	io.Copy(w, res.Body)
}
