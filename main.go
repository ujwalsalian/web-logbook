package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Web-Logbook!")
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
