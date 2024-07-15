package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/go-server/handlers" // Ensure this matches your module path
)

func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)
    fmt.Println("Server is running on port 4000")
    log.Fatal(http.ListenAndServe(":4000", nil))
}
