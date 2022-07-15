// net/http redirect - http.Redirect

package main

import (
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Request method at /index - "+r.Method+r.URL.Path)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Request method at "+r.URL.Path+" method "+r.Method)
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		io.WriteString(w, "Your url not found, redirect to \"/\"")
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Request method at /contact - "+r.Method)
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		io.WriteString(w, "Your url not found, redirect to \"/\"")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}
