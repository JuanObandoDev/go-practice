package main

import (
	"fmt"
	"net/http"
	"time"
)

// validae server status
func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println("Server", server, "is down!")
	} else {
		fmt.Println("Server", server, "is up and running!")
	}
}

func main() {
	// timer
	start := time.Now()

	// slice of servers
	servers := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
	}

	// loop through servers
	for _, server := range servers {
		checkServer(server)
	}
	finish := time.Since(start)
	fmt.Printf("Time taken: %s", finish)
}
