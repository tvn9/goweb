package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

type vehicle struct {
	Make      string
	Model     string
	YearBuilt string
}

func (v *vehicle) yearBuilt() {
	now := time.Now()
	v.YearBuilt = now.Format("01-02-2006")
}

func main() {
	f150 := vehicle{
		Make:  "Ford",
		Model: "F150",
	}
	f150.yearBuilt()
	tacomaSR := vehicle{
		Make:  "Toyota",
		Model: "Tacoma SR",
	}
	tacomaSR.yearBuilt()
	s150 := vehicle{
		Make:  "Chevrolet",
		Model: "Silverado 1500",
	}
	s150.yearBuilt()

	items := []vehicle{f150, tacomaSR, s150}

	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", items)
	if err != nil {
		log.Fatalln(err)
	}

}
