package main

import (
	"fmt"
	"net/http"
)

type port string

func (p port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Testing-Key", "this is the header test key")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var p port
	http.ListenAndServe(":8080", p)
}
