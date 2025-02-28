package main

import (
	"fmt"
	"net/http"
)

type page struct {
	name string
	id   int
}

type home struct {
	page
}

type contact struct {
	page
}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page.")
}

func (c *contact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the contact page.")
}

func main() {
	var h *home
	var c *contact

	http.Handle("/", h)
	http.Handle("/contact", c)
	http.ListenAndServe(":8080", nil)
}
