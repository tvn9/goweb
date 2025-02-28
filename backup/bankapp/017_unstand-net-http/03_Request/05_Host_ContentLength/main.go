package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
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
	tmpl.ExecuteTemplate(w, "index.html", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var t = Person{}
	http.ListenAndServe(":8080", &t)
}
