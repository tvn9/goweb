package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	str := "Thanh Pho Saigon."

	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", str)
	if err != nil {
		log.Fatalln(err)
	}
}
