package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

type Person struct {
	Fname      string
	Lname      string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.cio", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	err := tmpl.ExecuteTemplate(w, "index.html", Person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}
