package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/temmplate",
		Input:   `<script>alert("Hello World!");</script>`,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", home)
	if err != nil {
		log.Fatalln(err)
	}
}
