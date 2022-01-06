package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl *template.Template

// init
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

// Index
func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("fail to render template %v", err))
	}
}

// About
func about(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.tmpl", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("fail to render template %v", err))
	}
}
