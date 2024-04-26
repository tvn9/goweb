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

type person struct {
	Name string
	Job  string
}

func main() {
	thanh := person{
		Name: "Thanh",
		Job:  "Go Programmer",
	}

	err := tmpl.Execute(os.Stdout, thanh)
	if err != nil {
		log.Fatalln(err)
	}
}
