package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/tvn/goweb/utility"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/base.html"))
}

func main() {
	// Create a person
	p001 := utility.Person{
		Fname: "Mike",
		Lname: "McNillin",
		SNN:   utility.SSNCard{SSN: "009-99-9999", Country: "United States of America"},
		DL: utility.DLCard{ID: "A8352345", BD: "Jul-15-1975",
			State: "California", Expired: "08-12-2026", EyeColor: "Brown", Height: "5.5",
			Nationality: "Asian", Sex: "Male"},
		Phone: "555-555-9999",
		Email: "mmcnillin@email.com",
		Addr: utility.Address{StreetAddr: "555 St. Paul Blvd", City: "San Jose",
			State: "California", Zip: "93123", Country: "USA"},
	}
	// Open a new checking account
	ch002 := utility.Account{}
	ch002.Type = "Checking"
	ch002.Owner.Fname = "Thanh"
	ch002.Owner.Lname = "Nguyen"
	ch002.AccNum = "C000220200525"
	ch002.OpenBalance.Dollars = 0
	ch002.OpenBalance.Cents = 0
	ch002.Income.Dollars = 0
	ch002.Income.Cents = 0
	ch002.Expense.Dollars = 0
	ch002.Expense.Cents = 0
	ch002.EndBalance.Dollars = 0

	// Open the first account
	ch001 := utility.Account{
		Type:        "Checking",
		Owner:       p001,
		AccNum:      "C00120200525",
		OpenBalance: utility.Amount{Dollars: 0, Cents: 0},
		Income:      utility.Amount{Dollars: 0, Cents: 0},
		Expense:     utility.Amount{Dollars: 0, Cents: 0},
		EndBalance:  utility.Amount{Dollars: 0, Cents: 0},
	}

	// fmt.Println(ch001)

	// fmt.Println("Type: ", ch001.Type)
	// fmt.Println("Ownder:", ch001.Owner)
	// fmt.Println("Account No:", ch001.AccNum)
	// fmt.Println("OpenBalance:", ch001.OpenBalance.Dollars, ch001.OpenBalance.Cents)
	// fmt.Println("Income:", ch001.Income.Dollars, ch001.Income.Cents)
	fmt.Println("Transaction ID:", ch001.Activities.ID)

	ch001.Deposit(100, 50)

	ch001.WithDraw(50, 50)

	err := ExecuteTmpls("base.html", ch001)
	if err != nil {
		log.Fatal(err)
	}

	// err = ExecuteTmpls("base.html", ch002)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func ExecuteTmpls(str string, obj utility.Account) (err error) {
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, str, obj)
	if err != nil {
		return err
	}
	return err
}
