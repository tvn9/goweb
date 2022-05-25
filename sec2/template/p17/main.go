package main

import (
	"html/template"
	"log"
	"os"
)

type Course struct {
	Number, Name, Units string
}

type Semester struct {
	Term    string
	Courses []Cource
}

type Year struct {
	Fall, Spring, Summer Semester
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base.html"))
}

func main() {
	//
	y := Year{
		Fall: Semester{
			Term: "Fall",
			Courses: []Course{
				Course{"CSCI-40", "Introduction to Programming in Go", "4"},
				Course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				Course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: Semester{
			Term: "Spring",
			Courses: []Course{
				Course{"CSCI-50", "Advance Go", "5"},
				Course{"CSCI-190", "Advance Go", "5"},
			},
		},
	}
	err := tmpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}
}
