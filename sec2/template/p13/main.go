package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// Truck
type Truck struct {
	Owner     string
	Make      string
	Model     string
	Type      string
	YearBuilt int
	NoDoor    int
	TowCap    int
	LoadCap   int
	Speed     int
}

type Car struct {
	Owner     string
	Make      string
	Model     string
	Type      string
	YearBuilt int
	NoDoor    int
	Speed     int
	Color     string
	EPA       string
	Price     float64
}

type Vihicle struct {
	Trucks []Truck
	Cars   []Car
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	f150 := Truck{
		Owner:     "Thanh",
		Make:      "Ford Motors",
		Model:     "F1500",
		Type:      "Pickup Truck",
		YearBuilt: 2020,
		TowCap:    7500,
		LoadCap:   1500,
		Speed:     100,
	}

	sil150 := Truck{
		Owner:     "Xuan",
		Make:      "GM",
		Model:     "Silverado 1500",
		Type:      "Pickup Trick",
		YearBuilt: 2021,
		TowCap:    7500,
		LoadCap:   1500,
		Speed:     100,
	}

	sierra1500 := Truck{
		Owner:     "Bernie",
		Make:      "GMC",
		Model:     "Sierra 1500",
		Type:      "Pickup Truck",
		YearBuilt: 2022,
		TowCap:    9600,
		LoadCap:   2100,
		Speed:     100,
	}

	priusPrime := Car{
		Owner:     "Marko",
		Make:      "Toyota",
		Model:     "Prius Prime",
		Type:      "Saden",
		YearBuilt: 2022,
		NoDoor:    4,
		Speed:     120,
		Color:     "Red",
		EPA:       "55",
		Price:     28000.00,
	}

	rav4 := Car{
		Owner:     "Mike",
		Make:      "Toyota",
		Model:     "2022 RAV4 LE",
		Type:      "Mid SUV",
		YearBuilt: 2022,
		NoDoor:    4,
		Speed:     140,
		Color:     "Grey",
		EPA:       "30",
		Price:     26525.00,
	}

	vihGroup1 := Vihicle{
		Trucks: []Truck{f150, sil150, sierra1500},
		Cars:   []Car{priusPrime, rav4},
	}

	fn, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fn.Close()

	err = tmpl.ExecuteTemplate(os.Stdout, "base1.html", vihGroup1)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "base2.html", vihGroup1)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base1.html", vihGroup1)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(fn, "base2.html", vihGroup1)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range vihGroup1.Trucks {
		fmt.Printf("%s\n", t.Make)
		fmt.Printf("%s\n", t.Model)
		fmt.Printf("%d\n", t.YearBuilt)
		fmt.Printf("%d\n", t.TowCap)
		fmt.Printf("%d\n", t.LoadCap)
		fmt.Printf("%d\n", t.Speed)
	}

	for _, v := range vihGroup1.Cars {
		fmt.Printf("%s\n", v.Make)
		fmt.Printf("%s\n", v.Model)
		fmt.Printf("%d\n", v.YearBuilt)
		fmt.Printf("%s\n", v.Color)
		fmt.Printf("%s\n", v.EPA)
		fmt.Printf("%d\n", v.NoDoor)
		fmt.Printf("%d\n", v.Speed)
		fmt.Printf("%.2f\n", v.Price)
	}
}
