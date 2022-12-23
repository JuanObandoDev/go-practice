package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// middleware to authenticate
func CheckAuth() Middleware {
	// return middleware function
	return func(f http.HandlerFunc) http.HandlerFunc {
		// return handler function
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

// middleware to log
func Login() Middleware {
	// return middleware function
	return func(f http.HandlerFunc) http.HandlerFunc {
		// return handler function
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
