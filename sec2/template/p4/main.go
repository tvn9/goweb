// template.ParseFiles parsese multiple files example
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatal(err)
	}

	err = executeTmpl(tpl, "base1.html", "base2.html", "base3.html")
	if err != nil {
		log.Fatal(err)
	}
}

func executeTmpl(tmpl *template.Template, str ...string) (err error) {
	for _, f := range str {
		tmpl.ExecuteTemplate(os.Stdout, f, nil)
	}
	return err
}
