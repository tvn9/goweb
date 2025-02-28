package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "about.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "contact.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
