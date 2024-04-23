// Understanding net/http package
// example 4 - URL

package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type someThing struct{}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func (s someThing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("fails to parse form")
	}

	data := struct {
		Method     string
		URL        *url.URL
		Submission url.Values
		Header     http.Header
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	var s someThing
	http.ListenAndServe(":8080", s)
}
