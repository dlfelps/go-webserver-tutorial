package main

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
