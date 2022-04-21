// Passing data to html template using variable name
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/base1.html"))
}

func main() {

	// Get a string from command-line
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter some thing, example [\"The blue sky!\"]")
		return
	}

	str := args[0]
	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Fails to create file", err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", str)
	if err != nil {
		log.Fatal(err)
	}
}
