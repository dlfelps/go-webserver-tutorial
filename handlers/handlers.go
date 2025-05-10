package handlers

import (
        "fmt"
        "html/template"
        "net/http"
        "os"
        "path/filepath"
        "strings"
        "time"

        "golang-webserver-tutorial/content"
)

// TemplateData holds all data that will be passed to templates
type TemplateData struct {
        Title       string
        Content     template.HTML
        Tutorials   []content.Tutorial
        Examples    []content.CodeExample
        ActiveNav   string
        CurrentYear int
}

// parseTemplate parses the given template files and executes them with the provided data
func parseTemplate(w http.ResponseWriter, data TemplateData, templateFiles ...string) {
        // Add layout template to the list of templates
        files := append([]string{"templates/layout.html"}, templateFiles...)
        
        // Parse templates
        tmpl, err := template.ParseFiles(files...)
        if err != nil {
                http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
                return
        }
        
        // Execute template
        err = tmpl.ExecuteTemplate(w, "layout", data)
        if err != nil {
                http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
        }
}

// HomeHandler handles the root path and displays the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w, r)
                return
        }
        
        data := TemplateData{
                Title:       "Learn Go Web Development",
                ActiveNav:   "home",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/home.html")
}

// BasicHandler displays the basic concepts page
func BasicHandler(w http.ResponseWriter, r *http.Request) {
        tutorials := content.GetBasicTutorials()
        
        data := TemplateData{
                Title:       "Basic Web Server Concepts",
                Tutorials:   tutorials,
                ActiveNav:   "basic",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/basic.html")
}

// IntermediateHandler displays the intermediate concepts page
func IntermediateHandler(w http.ResponseWriter, r *http.Request) {
        tutorials := content.GetIntermediateTutorials()
        
        data := TemplateData{
                Title:       "Intermediate Web Server Concepts",
                Tutorials:   tutorials,
                ActiveNav:   "intermediate",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/intermediate.html")
}

// AdvancedHandler displays the advanced concepts page
func AdvancedHandler(w http.ResponseWriter, r *http.Request) {
        tutorials := content.GetAdvancedTutorials()
        
        data := TemplateData{
                Title:       "Advanced Web Server Concepts",
                Tutorials:   tutorials,
                ActiveNav:   "advanced",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/advanced.html")
}

// RestfulHandler displays the RESTful API concepts page
func RestfulHandler(w http.ResponseWriter, r *http.Request) {
        tutorials := content.GetRestfulTutorials()
        
        data := TemplateData{
                Title:       "RESTful API Development",
                Tutorials:   tutorials,
                ActiveNav:   "restful",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/restful.html")
}

// ExamplesHandler displays the code examples page
func ExamplesHandler(w http.ResponseWriter, r *http.Request) {
        examples := content.GetCodeExamples()
        
        data := TemplateData{
                Title:       "Code Examples",
                Examples:    examples,
                ActiveNav:   "examples",
                CurrentYear: time.Now().Year(),
        }
        
        parseTemplate(w, data, "templates/examples.html")
}

// DownloadHandler provides downloadable code examples
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
        // Extract the example file name from the URL
        examplePath := strings.TrimPrefix(r.URL.Path, "/download/")
        if examplePath == "" {
                http.Error(w, "No example specified", http.StatusBadRequest)
                return
        }
        
        // Validate the file path to prevent directory traversal
        filePath := filepath.Join("static", "examples", examplePath)
        fileInfo, err := os.Stat(filePath)
        if err != nil || fileInfo.IsDir() {
                http.NotFound(w, r)
                return
        }
        
        // Set appropriate headers for downloading
        w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(examplePath)))
        w.Header().Set("Content-Type", "text/plain")
        
        // Serve the file
        http.ServeFile(w, r, filePath)
}

// EnsureExamplesGenerated makes sure all example code files exist
func EnsureExamplesGenerated(examplesDir string) {
        // Create examples directory if it doesn't exist
        if _, err := os.Stat(examplesDir); os.IsNotExist(err) {
                os.MkdirAll(examplesDir, 0755)
        }
        
        // Generate all example files
        examples := content.GetCodeExamples()
        for _, example := range examples {
                filePath := filepath.Join(examplesDir, example.Filename)
                if _, err := os.Stat(filePath); os.IsNotExist(err) {
                        os.WriteFile(filePath, []byte(example.Code), 0644)
                }
        }
}
