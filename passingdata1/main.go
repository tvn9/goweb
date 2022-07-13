// Passing values through the URL

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("q")
	io.WriteString(w, "Searching: "+t)
}

// http://localhost:8080/?q=anything
