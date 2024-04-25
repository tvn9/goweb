package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
