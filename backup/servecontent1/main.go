// Serving Content - Using http.ServeContent

package main

import (
	"net/http"
	"os"
)

func pix(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fstat.ModTime(), f)
}

func main() {
	http.HandleFunc("/", pix)
	http.ListenAndServe(":8080", nil)
}
