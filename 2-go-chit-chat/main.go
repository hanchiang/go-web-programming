package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Handle routes
	mux.HandleFunc("/", index)

	// Start the server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
