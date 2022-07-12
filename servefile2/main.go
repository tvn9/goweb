// Serving File - using io.Copy
// Example 2

package main

import (
	"io"
	"net/http"
	"os"
)

func serveHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg" height=400></img> `)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	// read file
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", serveHtml)
	http.HandleFunc("/pix", serveFile)
	http.ListenAndServe(":8080", nil)
}
