package main

import (
	"html/template"
	"io"
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
	io.WriteString(w, "about.html")
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
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", mux)
}
