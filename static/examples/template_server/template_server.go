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
	
	// Define the template inline for simplicity
	tmpl := template.Must(template.New("page").Parse(`
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
	`))
	
	// Execute the template with the data
	tmpl.Execute(w, data)
}
