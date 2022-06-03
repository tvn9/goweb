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

type Person struct {
	Fname string
	Lname string
	Job   string
	ProLg []string
	Age   int
}

func main() {

	p1 := Person{
		Fname: "Thanh",
		Lname: "Ng",
		Job:   "Programmer",
		ProLg: []string{"html", "java script", "css", "go"},
		Age:   52,
	}

	fmt.Println(p1)

	// Get a string from command-line
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter some thing, example [\"The blue sky!\"]")
		return
	}

	str := args[0]
	fmt.Println(str)
	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Fails to create file", err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", p1)
	if err != nil {
		log.Fatal(err)
	}

}
