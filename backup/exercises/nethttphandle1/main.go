// Exercise - Create a website with three page (end points)
// and pass something data the the pages

package main

import (
	"log"
	"net/http"
	"text/template"
)

type CompanyInfo struct {
	Name    string
	Address string
	CEOName string
}

type CarInfo struct {
	Build      string
	Model      string
	Type       string
	MaxTrailer int
	MaxPayload int
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func (com *CompanyInfo) index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", com)
	if err != nil {
		log.Fatal("fails to execute template", err)
	}
}

func (com *CompanyInfo) about(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.html", com)
	if err != nil {
		log.Fatal("fails to execute template", err)
	}
}

func (car *CarInfo) car(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "car.html", car)
	if err != nil {
		log.Fatal("fails to execute template", err)
	}
}

func main() {
	comInfo := CompanyInfo{
		Name:    "World Best EV Builder",
		Address: "111 World St, Skype City, WS 99999",
		CEOName: "Mr. Car Man",
	}

	car1 := CarInfo{
		Build:      "World Best EV Builder",
		Model:      "1st EV Truck",
		Type:       "Pickup Truck",
		MaxTrailer: 12000,
		MaxPayload: 3000,
	}

	http.Handle("/", http.HandlerFunc(comInfo.index))
	http.Handle("/about", http.HandlerFunc(comInfo.about))
	http.Handle("/car", http.HandlerFunc(car1.car))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
