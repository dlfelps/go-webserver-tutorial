package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Apply middleware to handlers
	http.HandleFunc("/", loggerMiddleware(homeHandler))
	http.HandleFunc("/protected", loggerMiddleware(authMiddleware(protectedHandler)))
	http.HandleFunc("/public", loggerMiddleware(publicHandler))
	
	// Start the server
	fmt.Println("Server running at http://localhost:8080/")
	fmt.Println("Routes:")
	fmt.Println("  / - Home page")
	fmt.Println("  /protected - Protected page (requires auth header)")
	fmt.Println("  /public - Public page")
	fmt.Println("")
	fmt.Println("To access the protected page, add an 'Authorization: valid-token' header")
	
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Middleware for logging requests
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		// Call the next handler
		next(w, r)
		
		// Log the response time
		log.Printf("Completed in %v", time.Since(start))
	}
}

// Middleware for authentication
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for auth token (simplified example)
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized: no token provided", http.StatusUnauthorized)
			return
		}
		
		// In a real app, you'd validate the token here
		if token != "valid-token" {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		
		// If authenticated, proceed to the next handler
		next(w, r)
	}
}

// Handler functions

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
	fmt.Fprintln(w, "Try accessing /protected (with auth header) or /public")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a protected page.")
	fmt.Fprintln(w, "If you can see this, your authentication was successful.")
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a public page that anyone can access.")
}
