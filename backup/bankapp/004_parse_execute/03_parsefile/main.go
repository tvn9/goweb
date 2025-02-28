package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err = tmpl.ParseFiles("index.html", "about.html", "contact.html")
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

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
