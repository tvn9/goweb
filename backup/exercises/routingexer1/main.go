// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template
package main

import (
	"fmt"
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to home page!")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to dog page!")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to me page!")
}

const portNum = ":8080"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	fmt.Println("Starting server on port", portNum)
	http.ListenAndServe(portNum, nil)
}
