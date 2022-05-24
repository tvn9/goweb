// template.ParseFiles parsese multiple files example
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	// Must function is just a wrapper with error checking for template.ParseGlob
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func executeTmpl(tmpl *template.Template, str ...string) (err error) {
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, f := range str {
		tmpl.ExecuteTemplate(file, f, nil)
	}
	return err
}

func main() {
	err := executeTmpl(tmpl, "base1.html", "base2.html", "base3.html")
	if err != nil {
		log.Fatal(err)
	}
}
