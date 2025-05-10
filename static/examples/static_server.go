package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a file server that serves files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	
	// Handle all requests by serving a file of the same name
	http.Handle("/", fs)
	
	// Start the server
	fmt.Println("Static file server running at http://localhost:8080/")
	fmt.Println("Serving files from the ./static directory")
	http.ListenAndServe("localhost:8080", nil)
}
