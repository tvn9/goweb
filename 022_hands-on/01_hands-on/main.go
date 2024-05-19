package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page.")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the contact page.")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}
