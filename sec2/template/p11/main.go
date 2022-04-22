// Passing data from a go struct to html template
package main

import (
	"fmt"
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

type Trucks []Truck

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	f150 := Truck{
		Make:      "Ford Motors",
		Model:     "F1500",
		YearBuild: 2020,
		TowCap:    7500,
		LoadCap:   1500,
		Speed:     100,
	}

	sil150 := Truck{
		Make:      "GM",
		Model:     "Silverado 1500",
		YearBuild: 2021,
		TowCap:    7500,
		LoadCap:   1500,
		Speed:     100,
	}

	sierra1500 := Truck{
		Make:      "GMC",
		Model:     "Sierra 1500",
		YearBuild: 2022,
		TowCap:    9600,
		LoadCap:   2100,
		Speed:     100,
	}

	trucks := Trucks{f150, sil150, sierra1500}

	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", trucks)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "base2.html", trucks)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", trucks)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base2.html", trucks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trucks)
}
