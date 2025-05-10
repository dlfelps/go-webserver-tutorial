package content

import "html/template"

// CodeExample represents a downloadable code example
type CodeExample struct {
	Title       string
	Description template.HTML
	Filename    string
	Code        string
}

// GetCodeExamples returns all downloadable code examples
func GetCodeExamples() []CodeExample {
	return []CodeExample{
		{
			Title:       "Simple HTTP Server",
			Description: template.HTML("A basic HTTP server that responds with 'Hello, World!'"),
			Filename:    "simple_server.go",
			Code: `package main

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
`,
		},
		{
			Title:       "Static File Server",
			Description: template.HTML("A web server that serves static files from a directory"),
			Filename:    "static_server.go",
			Code: `package main

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
`,
		},
		{
			Title:       "HTML Template Server",
			Description: template.HTML("A server that renders HTML templates with dynamic data"),
			Filename:    "template_server.go",
			Code: `package main

import (
	"html/template"
	"net/http"
)

// PageData holds the data for our template
type PageData struct {
	Title   string
	Message string
	Items   []string
}

func main() {
	// Register the handler function
	http.HandleFunc("/", templateHandler)
	
	// Start the server
	http.ListenAndServe("localhost:8080", nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare the data
	data := PageData{
		Title:   "Template Demo",
		Message: "Welcome to Go Templates!",
		Items:   []string{"Item 1", "Item 2", "Item 3"},
	}
	
	// Define the template inline for simplicity
	tmpl := template.Must(template.New("page").Parse(` + "`" + `
		<!DOCTYPE html>
		<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<p>{{.Message}}</p>
			
			<ul>
				{{range .Items}}
					<li>{{.}}</li>
				{{end}}
			</ul>
		</body>
		</html>
	` + "`" + `))
	
	// Execute the template with the data
	tmpl.Execute(w, data)
}
`,
		},
		{
			Title:       "RESTful API Server",
			Description: template.HTML("A simple RESTful API server for a book collection"),
			Filename:    "rest_api.go",
			Code: `package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Book represents a book in our API
type Book struct {
	ID     int    ` + "`json:\"id\"`" + `
	Title  string ` + "`json:\"title\"`" + `
	Author string ` + "`json:\"author\"`" + `
	Year   int    ` + "`json:\"year\"`" + `
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
`,
		},
		{
			Title:       "Complete Web Application",
			Description: template.HTML("A more complete web application with routing, templates, and a mock database"),
			Filename:    "complete_app.go",
			Code: `package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Task represents a to-do item
type Task struct {
	ID     int
	Title  string
	Done   bool
	UserID int
}

// User represents a user of the application
type User struct {
	ID       int
	Username string
}

// PageData holds data for HTML templates
type PageData struct {
	Title    string
	User     *User
	Tasks    []Task
	TaskView *Task
	Flash    string
}

// In-memory "database"
var (
	users = []User{
		{ID: 1, Username: "alice"},
		{ID: 2, Username: "bob"},
	}
	
	tasks = []Task{
		{ID: 1, Title: "Learn Go basics", Done: true, UserID: 1},
		{ID: 2, Title: "Build a web server", Done: false, UserID: 1},
		{ID: 3, Title: "Deploy to production", Done: false, UserID: 1},
		{ID: 4, Title: "Study Go concurrency", Done: false, UserID: 2},
	}
	
	// Store our templates
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Register route handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	
	// Start the server
	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// indexHandler handles the home page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	// Get the current user
	userID := getUserIDFromCookie(r)
	var user *User
	if userID > 0 {
		user = getUserByID(userID)
	}
	
	// Prepare page data
	data := PageData{
		Title: "Task Manager",
		User:  user,
	}
	
	// Render the template
	templates.ExecuteTemplate(w, "index.html", data)
}

// tasksHandler handles the tasks list page
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in
	userID := getUserIDFromCookie(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	user := getUserByID(userID)
	
	// Get tasks for this user
	userTasks := getTasksByUserID(userID)
	
	// Handle form submission to add a task
	if r.Method == http.MethodPost {
		// Parse the form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		
		// Get the new task title
		title := r.FormValue("title")
		if title != "" {
			// Add the task
			addTask(title, userID)
			
			// Redirect to avoid form resubmission
			http.Redirect(w, r, "/tasks", http.StatusSeeOther)
			return
		}
	}
	
	// Prepare page data
	data := PageData{
		Title: "My Tasks",
		User:  user,
		Tasks: userTasks,
	}
	
	// Render the template
	templates.ExecuteTemplate(w, "tasks.html", data)
}

// taskHandler handles viewing and updating individual tasks
func taskHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in
	userID := getUserIDFromCookie(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	
	user := getUserByID(userID)
	
	// Extract the task ID from the URL
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	
	// Get the task
	task := getTaskByID(id)
	if task == nil {
		http.NotFound(w, r)
		return
	}
	
	// Check if the task belongs to the current user
	if task.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}
	
	// Handle task updates
	if r.Method == http.MethodPost {
		// Parse the form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		
		// Check if this is a delete request
		if r.FormValue("_method") == "DELETE" {
			// Delete the task
			deleteTask(id)
			
			// Redirect to the tasks list
			http.Redirect(w, r, "/tasks", http.StatusSeeOther)
			return
		}
		
		// Get form values
		title := r.FormValue("title")
		done := r.FormValue("done") == "on"
		
		// Update the task
		updateTask(id, title, done)
		
		// Redirect to the tasks list
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}
	
	// Prepare page data
	data := PageData{
		Title:    "Edit Task",
		User:     user,
		TaskView: task,
	}
	
	// Render the template
	templates.ExecuteTemplate(w, "task.html", data)
}

// loginHandler handles user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if already logged in
	if getUserIDFromCookie(r) > 0 {
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}
	
	// Initialize page data
	data := PageData{
		Title: "Login",
	}
	
	// Handle form submission
	if r.Method == http.MethodPost {
		// Parse the form
		if err := r.ParseForm(); err != nil {
			data.Flash = "Invalid form data"
			templates.ExecuteTemplate(w, "login.html", data)
			return
		}
		
		// Get the username
		username := r.FormValue("username")
		
		// Find the user
		user := getUserByUsername(username)
		if user == nil {
			data.Flash = "User not found"
			templates.ExecuteTemplate(w, "login.html", data)
			return
		}
		
		// Set a cookie to "log in" the user
		// In a real app, you'd use a secure session mechanism
		cookie := &http.Cookie{
			Name:     "user_id",
			Value:    strconv.Itoa(user.ID),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		
		// Redirect to the tasks page
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}
	
	// Render the login form
	templates.ExecuteTemplate(w, "login.html", data)
}

// logoutHandler handles user logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the user_id cookie
	cookie := &http.Cookie{
		Name:     "user_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	
	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Database helper functions

// getUserByID finds a user by ID
func getUserByID(id int) *User {
	for i := range users {
		if users[i].ID == id {
			return &users[i]
		}
	}
	return nil
}

// getUserByUsername finds a user by username
func getUserByUsername(username string) *User {
	for i := range users {
		if users[i].Username == username {
			return &users[i]
		}
	}
	return nil
}

// getTasksByUserID returns all tasks for a user
func getTasksByUserID(userID int) []Task {
	var userTasks []Task
	for _, task := range tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks
}

// getTaskByID finds a task by ID
func getTaskByID(id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

// addTask adds a new task
func addTask(title string, userID int) {
	// Find the highest task ID
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	
	// Create a new task
	newTask := Task{
		ID:     maxID + 1,
		Title:  title,
		Done:   false,
		UserID: userID,
	}
	
	// Add it to the tasks list
	tasks = append(tasks, newTask)
}

// updateTask updates an existing task
func updateTask(id int, title string, done bool) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = title
			tasks[i].Done = done
			break
		}
	}
}

// deleteTask removes a task
func deleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			// Remove the task by appending the tasks before and after it
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}

// getUserIDFromCookie gets the user ID from the cookie
func getUserIDFromCookie(r *http.Request) int {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return 0
	}
	
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0
	}
	
	return userID
}
`,
		},
		{
			Title:       "Middleware Example",
			Description: template.HTML("Example of creating and using middleware in Go web servers"),
			Filename:    "middleware.go",
			Code: `package main

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
`,
		},
	}
}
