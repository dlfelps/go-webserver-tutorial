package main

import (
        "fmt"
        "log"
        "net/http"
        "path/filepath"
        "time"

        "golang-webserver-tutorial/handlers"
)

func main() {
        // Define server port
        port := "5000"

        // Create a file server for static assets
        fs := http.FileServer(http.Dir("static"))
        http.Handle("/static/", http.StripPrefix("/static/", fs))

        // Register route handlers
        http.HandleFunc("/", handlers.HomeHandler)
        http.HandleFunc("/basic", handlers.BasicHandler)
        http.HandleFunc("/intermediate", handlers.IntermediateHandler)
        http.HandleFunc("/advanced", handlers.AdvancedHandler)
        http.HandleFunc("/restful", handlers.RestfulHandler)
        http.HandleFunc("/examples", handlers.ExamplesHandler)
        http.HandleFunc("/download/", handlers.DownloadHandler)

        // Create examples directory if it doesn't exist
        examplesDir := filepath.Join("static", "examples")
        handlers.EnsureExamplesGenerated(examplesDir)

        // Configure server
        server := &http.Server{
                Addr:           "0.0.0.0:" + port,
                ReadTimeout:    10 * time.Second,
                WriteTimeout:   10 * time.Second,
                MaxHeaderBytes: 1 << 20,
        }

        // Start the server
        fmt.Printf("Server running at http://0.0.0.0:%s/\n", port)
        log.Fatal(server.ListenAndServe())
}
