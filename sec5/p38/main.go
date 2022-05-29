package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home page route")
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "contact route")
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/contact", contact)

	// This server is using the default mux
	http.ListenAndServe(":8080", nil)
}
