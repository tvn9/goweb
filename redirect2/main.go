// Reditect - example http.StatusTemporaryRedirect

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal("fails to execute tamplete", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Printf("Requested url %s get redirecting to %s\n", r.URL.Path, " / ")
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Printf("Requested url %s get redirecting to %s\n", r.URL.Path, " / ")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}
