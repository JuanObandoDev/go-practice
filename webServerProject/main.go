package main

func main() {
	// start server in port 3000
	server := NewServer(":3000")

	// add route and handler function
	server.Handle("GET", "/", HandleRoot)

	// add route and handler function
	server.Handle("GET", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Login()))

	// server listening
	server.Listener()
}
