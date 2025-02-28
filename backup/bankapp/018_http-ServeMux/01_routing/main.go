package main

import (
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Name    string
	Address string
}

func (d *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "This is the home page.")
	case "/thanh":
		io.WriteString(w, "Thanh Vu Nguyen")
	case "/about":
		io.WriteString(w, "This is about page.")
	default:
		io.WriteString(w, "404 - Page not found.")
	}
}

func main() {
	var d *Data
	http.ListenAndServe(":8080", d)
}
