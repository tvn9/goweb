// Serving htmp go templates - using template.ParseGlob

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contact.html", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := tmpl.ExecuteTemplate(w, "applyProcess.html", nil)
		HandleError(w, err)
		return
	}

	err := tmpl.ExecuteTemplate(w, "apply.html", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
