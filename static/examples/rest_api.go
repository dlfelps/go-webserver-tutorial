package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Book represents a book in our API
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Sample data
var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan & Brian Kernighan", Year: 2015},
	{ID: 2, Title: "Go in Action", Author: "William Kennedy", Year: 2015},
	{ID: 3, Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Year: 2017},
}

func main() {
	// Register API routes
	http.HandleFunc("/api/books", booksHandler)
	http.HandleFunc("/api/books/", bookHandler) // For individual books
	
	// Start the server
	fmt.Println("RESTful API server running at http://localhost:8080/")
	http.ListenAndServe("localhost:8080", nil)
}

// booksHandler handles the collection of books
func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case http.MethodGet:
		// Return all books
		json.NewEncoder(w).Encode(books)
		
	case http.MethodPost:
		// Create a new book
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		// Set a new ID
		book.ID = len(books) + 1
		
		// Add to the collection
		books = append(books, book)
		
		// Return the created book with 201 Created status
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// bookHandler handles operations on individual books
func bookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract the book ID from the URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	
	// Find the book
	var index = -1
	var book Book
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
			break
		}
	}
	
	if index == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		// Return the book
		json.NewEncoder(w).Encode(book)
		
	case http.MethodPut:
		// Update the book
		var updatedBook Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		// Preserve the ID
		updatedBook.ID = id
		
		// Update the book
		books[index] = updatedBook
		
		// Return the updated book
		json.NewEncoder(w).Encode(updatedBook)
		
	case http.MethodDelete:
		// Remove the book
		books = append(books[:index], books[index+1:]...)
		
		// Return 204 No Content
		w.WriteHeader(http.StatusNoContent)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
