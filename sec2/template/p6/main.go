// Passing data to html template
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base1.html"))
}
func main() {
	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Fails to create file", err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", 42)
	if err != nil {
		log.Fatal(err)
	}

}
