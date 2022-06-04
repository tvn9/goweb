package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type port int

func (p port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		r.Method,
		r.URL,
		r.Form,
	}
	tmpl.ExecuteTemplate(w, "base.html", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base.html"))
}

func main() {
	var p port
	http.ListenAndServe(":8080", p)
}
