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

	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
