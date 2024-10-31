package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    // Set the port from the environment variable or use default
    port := os.Getenv("PORT")
    if port == "" {
        port = "4000" // Default port if not specified
    }

    // Serve static files from the current directory (which is app)
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/", fs) // Serve files from the current directory

    log.Printf("Server is running on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
