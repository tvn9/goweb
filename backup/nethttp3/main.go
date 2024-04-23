// Retrieving form values - exploring *http.Request
// example 2
package main

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"
)

type anything struct {
	// empty struct
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = errors.New("fails to parse form")
	}

	data := struct {
		Method     string
		Submission url.Values
	}{
		r.Method,
		r.Form,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	var any anything

	http.ListenAndServe(":8080", any)
}
