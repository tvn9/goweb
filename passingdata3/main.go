// Passing data - via Form or URL

package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	<input type="text" name="q">
	<input type="submit">
	</form>
	<br>`+v)
}

func form(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	err := tmpl.ExecuteTemplate(w, "index.html", v)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, v)
}

func main() {
	http.HandleFunc("/", form)
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
