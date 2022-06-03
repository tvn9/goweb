package main

import (
	"html/template"
	"log"
	"net/http"
)

func (d data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "base.html", r.Form)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base.html"))
}

type data string

func main() {
	var d data
	http.ListenAndServe(":8080", d)
}
