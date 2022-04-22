// Passing data from a go struct to html template
package main

import (
	"log"
	"os"
	"text/template"
)

// Truck
type Truck struct {
	Make      string
	Model     string
	YearBuild int
	TowCap    int
	LoadCap   int
	Speed     int
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	f150 := Truck{
		Make:      "Ford",
		Model:     "F1500",
		YearBuild: 2020,
		TowCap:    7500,
		LoadCap:   2500,
		Speed:     100,
	}
	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", f150)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "base2.html", f150)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", f150)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base2.html", f150)
	if err != nil {
		log.Fatal(err)
	}
}
