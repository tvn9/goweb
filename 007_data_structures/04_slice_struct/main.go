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
	t := person{
		Name: "Thanh",
		Job:  "Go Programmer",
	}
	x := person{
		Name: "Xuan",
		Job:  "Technical Writer",
	}

	employees := []person{t, x}

	err := tmpl.Execute(os.Stdout, employees)
	if err != nil {
		log.Fatalln(err)
	}
}
