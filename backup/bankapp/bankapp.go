package main

import (
	"bankapp/appcfg"
	"bankapp/calculator"
	"bankapp/database"
	"bankapp/reports"
	"errors"
	"fmt"
)

func main() {

	var userInput float64
	fmt.Println("Welcome to Bank App!")
	for userInput != 4 {

		printMainMenu()

		fmt.Print("You input: ")
		fmt.Scan(&userInput)

		if userInput == 1 {
			reports.TransactionReport(0, 0, database.ReadDataFromFile())
		} else if userInput == 2 {
			depositAmount, err := getUserInput("Enter deposit amount: ")
			if err != nil {
				fmt.Println("Error -", err)
			}
			calculator.Deposit(depositAmount)
		} else if userInput == 3 {
			withdrawAmount, err := getUserInput("Enter withdraw amount: ")
			if err != nil {
				fmt.Println("Error -", err)
			}
			calculator.Withdraw(withdrawAmount)
		} else if userInput == 4 {
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for banking with us!")
			return
		} else {
			fmt.Println("Invalid input!")
			fmt.Println()
		}
	}
}

func getUserInput(s string) (float64, error) {
	var value float64
	fmt.Print(s)
	fmt.Scan(&value)
	if value < 0 {
		return appcfg.ReturnZero, errors.New("Invalid amount, please ether positive amount!")
	} else {
		return value, nil
	}
}

// Exercises
// 1. Refacture main to move features inside main() to functions outside of main - Done
// 2. Refacture the code base to move functions from single file to pacakges - Done
// 3. Test the all features for errors and fix bugs. - Done
