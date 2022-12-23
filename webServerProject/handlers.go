package main

import (
	"fmt"
	"net/http"
)

// handler root  function
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// home handler function
func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}
