package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", srvRemoteFile)
	http.ListenAndServe(":8080", nil)
}

func srvRemoteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!-- posting image file from remote server -->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}
