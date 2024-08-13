package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to My Go Web App!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Println("Server started on: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}