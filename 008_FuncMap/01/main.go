package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

type person struct {
	Name string
	Job  string
}

type cloth struct {
	Name  string
	Model string
	Make  string
}

var fm = template.FuncMap{
	"fc": firstCapital,
	"ft": fisrtThree,
}

func init() {
	tmpl = template.Must(template.New("index.html").Funcs(fm).ParseFiles("index.html"))
}

func firstCapital(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func fisrtThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	t := person{
		Name: "thanh",
		Job:  "Go Programmer",
	}
	x := person{
		Name: "xuan",
		Job:  "Technical Writer",
	}
	b := person{
		Name: "bernie",
		Job:  "Product Manager",
	}

	ts := cloth{
		Name:  "MoaChachDong",
		Model: "L-Moa",
		Make:  "Mommy",
	}
	p := cloth{
		Name:  "TimeLessPeace",
		Model: "L-Time",
		Make:  "Mommy",
	}

	employees := []person{t, x, b}
	clothing := []cloth{ts, p}

	data := struct {
		Humen   []person
		Product []cloth
	}{
		Humen:   employees,
		Product: clothing,
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
