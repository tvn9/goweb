// Passing Data - Using Form text/html

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
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tmpl.ExecuteTemplate(w, "index.html", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
