package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at home: ", r.Method, "\n\n")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request method at about:", r.Method)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
