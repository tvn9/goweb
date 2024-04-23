// Understanding routing
// Exmaple 4 - HandlerFunc

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
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/about", http.HandlerFunc(about))

	http.ListenAndServe(":8080", nil)
}
