// Pasing data to html template and form

package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func index(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	p := Person{
		FirstName:  f,
		LastName:   l,
		Subscribed: s,
	}

	err := tmpl.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
