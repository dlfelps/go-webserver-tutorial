package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle all requests with the hello function
	http.HandleFunc("/", hello)
	
	// Start the server on port 8080
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe("localhost:8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Write a response to the client
	fmt.Fprintf(w, "Hello, World!")
}
