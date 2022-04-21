// template.ParseFiles parsese multiple files example
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	err = tmpl.ExecuteTemplate(nf, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(nf, "base2.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(nf, "base3.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "base2.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
