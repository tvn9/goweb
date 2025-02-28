package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Thanh struct {
	FirstName string
	LastName  string
}

func (t Thanh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var t Thanh
	http.ListenAndServe(":8000", t)
}
