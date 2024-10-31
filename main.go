package main

import (
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html") // Serve the HTML file
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "4000" // Default port if not specified
    }
    http.HandleFunc("/", handler)
    log.Printf("Server is running on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
