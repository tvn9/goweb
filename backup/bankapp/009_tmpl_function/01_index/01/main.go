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
	usPresidents := []string{"George Washington", "John Adams", "Thomas Jefferson", "James Madison"}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", usPresidents)
	if err != nil {
		log.Fatalln(err)
	}
}
