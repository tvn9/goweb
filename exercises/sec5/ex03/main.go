package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	Fname string
	Lname string
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.html"))
}

func main() {

	http.Handle("/", http.HandlerFunc(home))
	http.HandleFunc("/about", http.HandlerFunc(about))

	http.ListenAndServe(":8080", nil)
}
