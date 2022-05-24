// template.ParseFiles example
package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func main() {
	tmpl, err := tmpl.ParseFiles("base.html")
	if err != nil {
		log.Fatal(err)
	}
	// let parse some more templates
	tmpl, err = tmpl.ParseFiles("base1.html", "base2.html", "base3.html")
	if err != nil {
		log.Fatal(err)
	}

	// Create a file name index.html
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Let execute a specific template to terminal
	errCheck(executeTmpl(tmpl, "base.html", "base1.html", "base2.html", "base3.html"))

	err = tmpl.Execute(file, tmpl)
	if err != nil {
		log.Fatal(err)
	}
}

func executeTmpl(tmpl *template.Template, str ...string) (err error) {
	for _, f := range str {
		err = tmpl.ExecuteTemplate(os.Stdout, f, nil)
	}
	return err
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
