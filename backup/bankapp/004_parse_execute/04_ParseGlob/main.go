package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
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
