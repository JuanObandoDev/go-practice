package main

import "net/http"

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// User type
type User struct {
	Name  string
	Email string
	Phone string
}

// MetaData type
type MetaData interface{}
