package main

import (
	"fmt"
	"net/http"
	"time"
)

// validate server status
func checkServer(server string, ch chan string) {
	_, err := http.Get(server)
	if err != nil {
		// adding data to channel
		ch <- server + " is down!"
	} else {
		// adding data to channel
		ch <- server + " is up and running!"
	}
}

func main() {
	// timer
	start := time.Now()

	// channels
	ch := make(chan string)

	// slice of servers
	servers := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
	}

	// loop through servers
	for _, server := range servers {
		// implementing goroutine
		go checkServer(server, ch)

		// receiving data from channel
		fmt.Println(<-ch)
	}
	finish := time.Since(start)
	fmt.Printf("Time taken: %s", finish)
}
