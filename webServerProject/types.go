package main

import "net/http"

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc
