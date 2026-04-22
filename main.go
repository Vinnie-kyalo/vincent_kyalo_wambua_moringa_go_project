package main

import (
	"fmt"
	"net/http"
)

// Home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go Web Server!")
}

// JSON status route
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","service":"go_project"}`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}