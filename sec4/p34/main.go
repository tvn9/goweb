package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type port string

func (p port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
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
