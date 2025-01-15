package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from MondooEngineer!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on port 4733...")
    if err := http.ListenAndServe(":4733", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
