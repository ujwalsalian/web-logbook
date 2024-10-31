package main

import (
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, web-logbook!"))
    })

    log.Println("Server is running on port 4000...")
    err := http.ListenAndServe(":4000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
