package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", "Thanh Nguyen")
	if err != nil {
		log.Fatalln(err)
	}
}
