// Retrieving form values - exploring *http.Request
package main

import (
	"errors"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

type anything struct{}

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = errors.New("fail to parse form")
	}
	tmpl.ExecuteTemplate(w, "index.html", r.Form)
}

func main() {
	var any anything

	http.ListenAndServe(":8080", any)
}
