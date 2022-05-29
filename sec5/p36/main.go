package main

import (
	"io"
	"net/http"
)

type port string

func (p port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "This is the home page using the /")
	case "/login":
		io.WriteString(w, "The login page")
	case "/contact":
		io.WriteString(w, "The contact page")
	default:
		io.WriteString(w, "Sorry the page is not exist")
	}
}

func main() {
	var p port
	http.ListenAndServe(":8080", p)
}
