package main

import (
	"encoding/json"
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

// post request handler function
func PostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Fprintf(w, "Metadata: %v", metadata)
}

// userPostRequest handler function
func UserPostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	res, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
