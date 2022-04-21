// template.ParseFiles parsese multiple files example
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("base1.html", "base2.html", "base3.html")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(nf, "base2.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(nf, "base3.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "base1.html", nil)
	if err != nil {
		log.Fatal(err)
	}

}
