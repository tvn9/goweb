package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", laptop)
	http.HandleFunc("/mac1.jpeg", macPic)

	http.ListenAndServe(":8080", nil)
}

func laptop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/mac1.jpeg">`)
}

func macPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("mac1.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
