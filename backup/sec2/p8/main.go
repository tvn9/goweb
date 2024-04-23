// Passing a slice of data into html template
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
	str := []string{"Connecticut", "Delaware", "Georgia", "Maryland", "Massachusetts",
		"New Jersey", "New York", "New Hamshire", "North Carolina", "Pennsylvania",
		"Rhode Island", "South Carolina", "Virginia"}

	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("fails to create file", err)
	}
	defer fn.Close()

	err = tmpl.Execute(os.Stdout, str)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", str)

}
