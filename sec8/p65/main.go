package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at home", r.Method, "\n\n")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at about", r.Method)
	// process form data
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You request method at about", r.Method)
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
