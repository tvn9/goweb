package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Person struct {
	FirstName string
	LastLast  string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
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
	tmpl.ExecuteTemplate(w, "index.html", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	thanh := &Person{
		FirstName: "Thanh",
		LastLast:  "Nguyen",
	}
	http.ListenAndServe(":8080", thanh)
}
