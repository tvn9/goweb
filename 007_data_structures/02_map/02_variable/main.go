package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	// sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
