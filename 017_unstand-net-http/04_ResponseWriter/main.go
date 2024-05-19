package main

import (
	"fmt"
	"net/http"
)

type Name string

func (n Name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("First-Key", "Thanh Vu Nguyen")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Testing</h1>")
}

func main() {
	var t Name
	http.ListenAndServe(":8080", t)
}
