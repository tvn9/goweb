// Serving Files - with http.FileServer

package main

import (
	"io"
	"net/http"
)

func pix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/pix", pix)
	http.ListenAndServe(":8080", nil)
}
