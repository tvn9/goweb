package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", srvFile)
	http.HandleFunc("/code1.jpeg", codeImage)
	http.ListenAndServe(":8080", nil)
}

func srvFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="code1.jpeg">`)
}

func codeImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "code1.jpeg")
}
