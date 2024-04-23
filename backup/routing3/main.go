// Understanding Routing
// Example 3 - Default ServeMux

package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to home page!")
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to about page!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}
