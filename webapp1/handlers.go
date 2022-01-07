package main

import (
	"log"
	"net/http"
	"text/template"
)

// tmpl
var tmpl *template.Template

// init
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

// index
func Index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.tmpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}
