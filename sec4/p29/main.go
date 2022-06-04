package main

import (
	"fmt"
	"net/http"
)

type portNum string

func (p portNum) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello HTTP Server!, Type Can be any thing")
}

func main() {
	var p portNum
	http.ListenAndServe(":8080", p)
}
