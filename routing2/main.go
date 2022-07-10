// Understanding routing
// Routing example 2 - http.NewServeMux

package main

import (
	"net/http"
)

type home struct{}
type about struct{}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to homepage!</h1>"))
}

func (a about) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to about page!</h1>"))
}

func main() {
	var h home
	var a about

	mux := http.NewServeMux()
	mux.Handle("/home", h)
	mux.Handle("/", h)
	mux.Handle("/about", a)

	http.ListenAndServe(":8080", mux)
}
