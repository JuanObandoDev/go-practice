package main

import (
	"net/http"
)

// router struct using map to store routes and handlers functions
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// router constructor
func NewRouter() *Router {
	return &Router{rules: make(map[string]map[string]http.HandlerFunc)}
}

// method to add routes and handlers functions
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, okPath, okMethod := r.FindHandlder(req.URL.Path, req.Method)
	if !okPath {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !okMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, req)
}

// method to find handler function
func (r *Router) FindHandlder(path string, method string) (http.HandlerFunc, bool, bool) {
	_, okPath := r.rules[path]
	handler, okMethod := r.rules[path][method]
	return handler, okPath, okMethod
}
