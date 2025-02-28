package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("index.html").Funcs(fm).ParseFiles("index.html"))
}

var fm = template.FuncMap{
	"formatMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())
	if err != nil {
		log.Fatalln("fails to format time, error:", err)
	}
}
