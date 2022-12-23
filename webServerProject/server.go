package main

import "net/http"

// server struct using router struct
type Server struct {
	port   string
	router *Router
}

// server constructor
func NewServer(port string) *Server {
	return &Server{
		port: port,
		// starting router
		router: NewRouter(),
	}
}

// method to server listening
func (s *Server) Listener() error {
	// using router as handler
	http.Handle("/", s.router)

	// listening in port 3000
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

// method to add routes and handlers functions
func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, ok := s.router.rules[path]
	if !ok {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// method to add middleware
func (s *Server) AddMiddleware(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	for _, middleware := range m {
		f = middleware(f)
	}
	return f
}
