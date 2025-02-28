package main

import (
	"html/template"
	"log"
	"net/http"
)

type Thanh struct {
	First string
	Last  string
}

func (m Thanh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "index.html", r.Form)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var t Thanh
	http.ListenAndServe(":8080", t)
}
