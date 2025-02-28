package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("index.html").Funcs(fm).ParseFiles("index.html"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
