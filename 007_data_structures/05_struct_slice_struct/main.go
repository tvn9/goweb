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

type cloth struct {
	Name  string
	Model string
	Make  string
}

type items struct {
	Humen   []person
	Product []cloth
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
	b := person{
		Name: "Bernie",
		Job:  "Product Manager",
	}

	ts := cloth{
		Name:  "Moa",
		Model: "L-Moa",
		Make:  "Mommy",
	}
	p := cloth{
		Name:  "Time",
		Model: "L-Time",
		Make:  "Mommy",
	}

	employees := []person{t, x, b}
	clothing := []cloth{ts, p}

	data := items{
		Humen:   employees,
		Product: clothing,
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
