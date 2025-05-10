package content

import (
	"html/template"
)

// Tutorial represents a single tutorial with title, description, and code examples
type Tutorial struct {
	ID          string
	Title       string
	Description template.HTML
	Code        template.HTML
	Explanation template.HTML
}

// GetBasicTutorials returns all basic level tutorials
func GetBasicTutorials() []Tutorial {
	return []Tutorial{
		{
			ID:    "hello-world",
			Title: "Hello World Web Server",
			Description: template.HTML(`
				<p>This is the simplest possible web server in Go. It responds with "Hello, World!" to every request.</p>
				<p>The <code>net/http</code> package provides all the functionality needed to create HTTP servers and clients.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
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
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>How It Works:</h4>
				<ul>
					<li><code>http.HandleFunc("/")</code> registers a function to handle all requests to the root path.</li>
					<li><code>http.ListenAndServe</code> starts an HTTP server listening on the specified address.</li>
					<li>The second parameter to <code>ListenAndServe</code> is a handler. <code>nil</code> means use the default router.</li>
					<li>Our <code>hello</code> function gets the <code>http.ResponseWriter</code> and <code>http.Request</code> parameters.</li>
					<li>Using <code>fmt.Fprintf</code>, we write our response text to the response writer.</li>
				</ul>
			`),
		},
		{
			ID:    "serve-html",
			Title: "Serving HTML Pages",
			Description: template.HTML(`
				<p>Most web servers need to serve HTML pages. Here's how to serve static HTML content in Go.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
package main

import (
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Handle the home page
	http.HandleFunc("/", homePage)
	
	// Start the server
	http.ListenAndServe("localhost:8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Serve the home page HTML file
	http.ServeFile(w, r, "templates/index.html")
}
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>How It Works:</h4>
				<ul>
					<li><code>http.FileServer</code> creates a handler that serves files from the given directory.</li>
					<li><code>http.StripPrefix</code> removes the given prefix from the URL path before passing it to the handler.</li>
					<li><code>http.ServeFile</code> serves a specific file in response to a request.</li>
					<li>Static files (CSS, JavaScript, images) are served from the "static" directory.</li>
					<li>HTML templates are served from the "templates" directory.</li>
				</ul>
			`),
		},
		{
			ID:    "handling-routes",
			Title: "Handling Different URL Routes",
			Description: template.HTML(`
				<p>A web server needs to handle different routes (URLs) differently. Here's how to implement basic routing in Go.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register handlers for different routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	
	// Start the server
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure we're at the root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the Home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us page")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Us page")
}
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>How It Works:</h4>
				<ul>
					<li>We register different handler functions for different URL paths using <code>http.HandleFunc</code>.</li>
					<li>Each handler function can perform different actions based on the route.</li>
					<li>In the <code>homeHandler</code>, we check if the path is exactly "/" and return a 404 error if not.</li>
					<li>This is important because the "/" route matches all paths that don't match other routes.</li>
					<li>For more complex routing, consider using router libraries like Gorilla Mux or Chi.</li>
				</ul>
			`),
		},
	}
}

// GetIntermediateTutorials returns all intermediate level tutorials
func GetIntermediateTutorials() []Tutorial {
	return []Tutorial{
		{
			ID:    "html-templates",
			Title: "Using HTML Templates",
			Description: template.HTML(`
				<p>Go's <code>html/template</code> package provides a powerful way to create dynamic HTML pages.</p>
				<p>It allows you to insert dynamic content into HTML templates, with automatic HTML escaping to prevent XSS attacks.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
package main

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
	
	// Parse the template file
	tmpl, err := template.ParseFiles("templates/demo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>How It Works:</h4>
				<ul>
					<li><code>template.ParseFiles</code> loads and parses the template file.</li>
					<li><code>tmpl.Execute</code> fills in the template with the provided data and writes to the response writer.</li>
					<li>In the template file, <code>{{.FieldName}}</code> inserts the value of the field.</li>
					<li><code>{{range .Items}}</code> loops over the Items slice.</li>
					<li>Go templates automatically escape HTML to prevent XSS attacks.</li>
					<li>The <code>html/template</code> package handles nested templates, conditionals, and more.</li>
				</ul>
			`),
		},
	}
}

// GetAdvancedTutorials returns all advanced level tutorials
func GetAdvancedTutorials() []Tutorial {
	return []Tutorial{
		{
			ID:    "json-apis",
			Title: "Building JSON APIs",
			Description: template.HTML(`
				<p>Go has excellent support for working with JSON, making it easy to build JSON APIs.</p>
				<p>Let's explore how to create JSON endpoints, handle JSON requests, and parse JSON data.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a user in our system
type User struct {
	ID       int    ` + "`json:\"id\"`" + `
	Username string ` + "`json:\"username\"`" + `
	Email    string ` + "`json:\"email\"`" + `
}

// Simple in-memory database
var users = []User{
	{ID: 1, Username: "johndoe", Email: "john@example.com"},
	{ID: 2, Username: "janedoe", Email: "jane@example.com"},
}

func main() {
	// API endpoints
	http.HandleFunc("/api/users", usersHandler)
	
	// Start the server
	fmt.Println("JSON API server running at http://localhost:8080/")
	http.ListenAndServe("localhost:8080", nil)
}

// usersHandler handles the collection of users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Return all users as JSON
	json.NewEncoder(w).Encode(users)
}
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>How It Works:</h4>
				<ul>
					<li><code>encoding/json</code> package provides functions for working with JSON data.</li>
					<li>The <code>json:\"field_name\"</code> struct tags tell the encoder what to name fields in the JSON output.</li>
					<li><code>json.NewEncoder(w).Encode(data)</code> writes JSON data to the response writer.</li>
					<li>We set <code>Content-Type: application/json</code> in the response headers.</li>
				</ul>
			`),
		},
	}
}

// GetRestfulTutorials returns all RESTful API tutorials
func GetRestfulTutorials() []Tutorial {
	return []Tutorial{
		{
			ID:    "rest-basics",
			Title: "RESTful API Basics",
			Description: template.HTML(`
				<p>REST (Representational State Transfer) is an architectural style for designing networked applications.</p>
				<p>RESTful APIs use HTTP methods explicitly and are stateless, with resources identified by URLs.</p>
			`),
			Code: template.HTML(`
<pre><code class="language-go">
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product represents a product in our API
type Product struct {
	ID          int     ` + "`json:\"id\"`" + `
	Name        string  ` + "`json:\"name\"`" + `
	Description string  ` + "`json:\"description\"`" + `
	Price       float64 ` + "`json:\"price\"`" + `
	Category    string  ` + "`json:\"category\"`" + `
}

// In-memory product database
var products = []Product{
	{ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 1299.99, Category: "Electronics"},
	{ID: 2, Name: "Headphones", Description: "Noise-cancelling headphones", Price: 249.99, Category: "Electronics"},
	{ID: 3, Name: "Coffee Maker", Description: "Automatic coffee maker", Price: 89.99, Category: "Kitchen"},
}

func main() {
	// Register API endpoints
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)
	
	// Start the server
	fmt.Println("RESTful API server running at http://localhost:8080/")
	http.ListenAndServe("localhost:8080", nil)
}

// productsHandler handles the collection endpoint
func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Return all products
	json.NewEncoder(w).Encode(products)
}

// productHandler handles the single-resource endpoint
func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract the product ID from the URL
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	
	// Find the product
	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	
	http.NotFound(w, r)
}
</code></pre>
			`),
			Explanation: template.HTML(`
				<h4>RESTful Principles:</h4>
				<ul>
					<li><strong>Resource-Based:</strong> Everything is a resource, identified by a URL (/products, /products/1)</li>
					<li><strong>HTTP Methods:</strong> Use standard HTTP methods for operations (GET, POST, PUT, DELETE)</li>
					<li><strong>Stateless:</strong> Each request contains all information needed to process it</li>
					<li><strong>Status Codes:</strong> Use appropriate HTTP status codes (200 OK, 404 Not Found, etc.)</li>
				</ul>
			`),
		},
	}
}