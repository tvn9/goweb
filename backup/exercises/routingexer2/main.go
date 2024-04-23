// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/"
// "/dog/"
// "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal("fails to execute template")
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "dog.html", nil)
	if err != nil {
		log.Fatal("fails to execute tempalte")
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "me.html", nil)
	if err != nil {
		log.Fatal("fails to execute tempalte")
	}
}

const portNum = ":8080"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	fmt.Println("Starting server on port", portNum)
	log.Fatal(http.ListenAndServe(portNum, nil))

}
