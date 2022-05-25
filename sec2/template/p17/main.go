package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Course struct {
	Number, Name, Units string
}

type Semester struct {
	Term    string
	Courses []Course
}

type Year struct {
	Year                 int
	Fall, Spring, Summer Semester
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base.html"))
}

func main() {
	//
	y := Year{
		Year: 2020,
		Fall: Semester{
			Term: "Fall",
			Courses: []Course{
				{"CSCI-40", "Introduction to Programming in Go", "4"},
				{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: Semester{
			Term: "Spring",
			Courses: []Course{
				{"CSCI-50", "Advance Go", "5"},
				{"CSCI-190", "Advance Web Programming with Go", "5"},
				{"CSCI-191", "Advance Mobile Apps With Go", "5"},
			},
		},
	}

	fmt.Println(y)
	err := tmpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, "base.html", y)
	if err != nil {
		log.Fatal()
	}

}
