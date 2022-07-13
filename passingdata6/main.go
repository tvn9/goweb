// Passing data - Using r.FormFile, io.ReadAll and os.Create

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	var txt string
	if r.Method == http.MethodPost {

		// Open a file
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Printf("file: %s\n, header: %v\n, err: %s\n", f, h, err)

		// Read data from file
		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		txt = string(bs)

		// send file to server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tmpl.ExecuteTemplate(w, "index.html", txt)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
