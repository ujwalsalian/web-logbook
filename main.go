package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", HomeHandler) // Ensure this points to your main handler
    // Add more routes as needed

    // Start the server
    log.Fatal(http.ListenAndServe(":4000", nil)) // Ensure the port matches
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "path/to/index.html") // Change this to your actual file path
}
