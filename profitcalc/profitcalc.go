package main

import (
	"errors"
	"fmt"
	"os"
)

const datafile = "datafile.txt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxrate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxrate)

	revenue, err := getUserInput("Enter revenue: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	expenses, err := getUserInput("Enter expenses: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	taxrate, err := getUserInput("Enter tax rate: ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// ebt := revenue - expenses
	// taxrate = taxrate / 100
	// profit := ebt - ebt*taxrate
	// ratio := ebt / profit

	ebt, taxrate, profit, ratio := profitCalc(revenue, expenses, taxrate)

	writeDataToFile(revenue, expenses, ebt, taxrate, profit, ratio)

	fmt.Printf("Revenue: %.2f\n", revenue)
	fmt.Printf("Expenses: %.2f\n", expenses)
	fmt.Printf("Tax rate = %.2f%%\n", taxrate*100)
	fmt.Println("-------------------------------")
	fmt.Printf("Earning before tax: %.2f\n", ebt)
	fmt.Printf("Profit after tax: %.2f\n", profit)
	fmt.Printf("Earning over profit ratio: %.2f\n", ratio)
}

// Refactor the above code with functions

func getUserInput(s string) (float64, error) {
	var value float64
	fmt.Print(s)
	fmt.Scan(&value)
	if value < 0 {
		return value, errors.New("Invalid amount, must be larger than 0")
	} else {
		return value, nil
	}
}

func profitCalc(rev, exp, tax float64) (ebt, taxrate, profit, ratio float64) {
	ebt = rev - exp
	taxrate = tax / 100
	profit = ebt - ebt*taxrate
	ratio = ebt / profit
	return ebt, taxrate, profit, ratio
}

func writeDataToFile(rev, exp, ebt, taxrate, profit, ratio float64) {
	data := fmt.Sprintf(`Profit & Loss Report
	Revenue: %.2f
	Expenses: %.2f
	Tax rate: %.2f
	Earning before tax: %.2f
	Profit: %.2f
	Earning over profit ratio: %.2f`, rev, exp, taxrate, ebt, profit, ratio)
	if err := os.WriteFile(datafile, []byte(data), 0644); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
