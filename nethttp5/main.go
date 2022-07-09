// Understanding net/http package
// Exmaple 5 - Header

package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Car struct{}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func (c Car) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	var c Car
	http.ListenAndServe(":8080", c)
}
