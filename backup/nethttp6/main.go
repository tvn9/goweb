// Understanding net/http package
// Example 6 - r.Host, r.ContentLength

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
		log.Fatal(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submission    map[string][]string
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

func main() {
	var s someThing

	http.ListenAndServe(":8080", s)
}
