package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/go-server/handlers"
)

func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
