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
	Fname      string
	Lname      string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", loadPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func loadPage(w http.ResponseWriter, r *http.Request) {
	// body
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tmpl.ExecuteTemplate(w, "index.html", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}
