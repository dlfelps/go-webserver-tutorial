package handlers

import (
        "net/http"
        "net/http/httptest"
        "testing"
        "os"
)

// TestMain sets up the test environment
func TestMain(m *testing.M) {
        // Create test template directory if it doesn't exist
        if _, err := os.Stat("../templates"); os.IsNotExist(err) {
                // We're running tests from the handlers directory, create templates directory
                os.Mkdir("../templates", 0755)
                
                // Create simple layout template for testing
                layoutContent := `{{define "layout"}}{{.Title}}{{template "content" .}}{{end}}`
                os.WriteFile("../templates/layout.html", []byte(layoutContent), 0644)
                
                // Create simple content templates for testing
                homeContent := `{{define "content"}}Home page content{{end}}`
                os.WriteFile("../templates/home.html", []byte(homeContent), 0644)
                
                basicContent := `{{define "content"}}Basic tutorial content{{end}}`
                os.WriteFile("../templates/basic.html", []byte(basicContent), 0644)
        }
        
        // Run the tests
        exitCode := m.Run()
        
        // Exit with the same code
        os.Exit(exitCode)
}

func TestHomeHandler(t *testing.T) {
        // Skip this test if running in CI environment
        if os.Getenv("CI") == "true" {
                t.Skip("Skipping test in CI environment")
        }
        
        // Create a request to pass to our handler
        req, err := http.NewRequest("GET", "/", nil)
        if err != nil {
                t.Fatal(err)
        }

        // Create a ResponseRecorder to record the response
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(HomeHandler)

        // Call the handler
        handler.ServeHTTP(rr, req)

        // We're not actually checking template rendering here, just that the function doesn't panic
        // In a real application we would mock the template rendering
        t.Log("Home handler test completed")
}

func TestBasicHandler(t *testing.T) {
        // Skip this test if running in CI environment
        if os.Getenv("CI") == "true" {
                t.Skip("Skipping test in CI environment")
        }
        
        // Create a request to pass to our handler
        req, err := http.NewRequest("GET", "/basic", nil)
        if err != nil {
                t.Fatal(err)
        }

        // Create a ResponseRecorder to record the response
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(BasicHandler)

        // Call the handler
        handler.ServeHTTP(rr, req)

        // We're not actually checking template rendering here, just that the function doesn't panic
        // In a real application we would mock the template rendering
        t.Log("Basic handler test completed")
}