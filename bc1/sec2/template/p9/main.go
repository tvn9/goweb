// Passing a slice of data into html template
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base1.html", "base2.html"))
}

func main() {
	str := map[string]string{
		"New York":  "Albany",
		"Maryland":  "Annapolis",
		"Georgia":   "Atlanta",
		"Maine":     "Augusta",
		"Texas":     "Austin",
		"Louisiana": "Baton Rouge",
	}

	fn1, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("fails to create file", err)
	}
	defer fn1.Close()

	fn2, err := os.Create("index2.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fn2.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", str)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "base2.html", str)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn1, "base1.html", str)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn2, "base2.html", str)
	if err != nil {
		log.Fatal(err)
	}
}
