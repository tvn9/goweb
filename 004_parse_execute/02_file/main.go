package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, nil)
	if err != nil {
		log.Fatal(err)
	}
}
