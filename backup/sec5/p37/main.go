package main

import (
	"io"
	"log"
	"net/http"
)

type home int

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the home page")
}

type contact string

func (c contact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This the contact route")
}

type login int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the login rout")
}

func main() {
	var h home
	var c contact
	var l login

	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/contact", c)
	mux.Handle("/login", l)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
