package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.html"))
}

type vehicle struct {
	Make      string
	Model     string
	YearBuilt string
}

func new() vehicle {
	v := vehicle{}
	t := time.Now()
	v.YearBuilt = t.Format("01-02-2006")
	return v
}

func (v *vehicle) insertModel(m string) {
	v.Model = m
}

func (v *vehicle) insertMake(m string) {
	v.Make = m
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	fmt.Println("starting server on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	f150 := new()
	tacomasr := new()

	f150.insertMake("Ford")
	f150.insertModel("F150")

	tacomasr.insertMake("Toyota")
	tacomasr.insertModel("Tacoma SR")

	items := []vehicle{f150, tacomasr}
	err := tmpl.ExecuteTemplate(w, "index.html", items)
	if err != nil {
		log.Fatalln("fail to execute template - error:", err)
	}
}
