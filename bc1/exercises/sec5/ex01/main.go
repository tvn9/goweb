package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dog Dog Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thanh Nguyen")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
