package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home page route")
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Contact page route")
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/contact", http.HandlerFunc(contact))

	http.ListenAndServe(":8080", nil)
}
