// Nested Tamplates
package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	file, err := os.Create("home.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(file, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
