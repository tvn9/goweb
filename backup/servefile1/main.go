// Serving File - using io.WriteString
// Example 1

package main

import (
	"io"
	"net/http"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	// Set content as html for source file
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(w, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	io.WriteString(w, `<img src="toby.jpg">`)
}

func main() {
	http.HandleFunc("/", serveFile)
	http.ListenAndServe(":8080", nil)
}
