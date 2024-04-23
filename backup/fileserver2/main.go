// Serving Files - Using http.FileServer

package main

import (
	"io"
	"net/http"
)

func pix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="./assets/toby.jpg">`)
}

func main() {
	http.HandleFunc("/", pix)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
