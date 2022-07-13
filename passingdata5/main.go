// passing data using io.ReadAll

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("textArea")

	log.Println(r.Method)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		log.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t = string(bs)
	}

	err := tmpl.ExecuteTemplate(w, "index.html", t)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
