package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("base.html"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "Ian Fleming",
			Age:  56,
		},
		false,
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(file, "base.html", p1)
	if err != nil {
		log.Fatal(err)
	}

}
