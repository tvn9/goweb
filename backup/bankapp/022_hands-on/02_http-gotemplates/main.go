package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", "Thanh")
	if err != nil {
		log.Fatal(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "about.html")
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	// mux := http.NewServeMux()

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}
