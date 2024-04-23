// Net.http package - htto.NotFoundHandler
// Example 1

package main

import (
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		io.WriteString(w, r.URL.Path)
		return
	}

	io.WriteString(w, r.URL.Path)
	fmt.Println(r.URL.Path)
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
