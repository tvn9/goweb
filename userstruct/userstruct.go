package main

import (
	"fmt"
	"strconv"
	"userstruct/inputs"
	"userstruct/users"
)

// Practice the struct data type and passing pointer to function

func main() {

	// Declare and blank user
	var u users.User

	// Print the app welcome message once
	fmt.Println("Welcome to user app!")

	for {
		// Print the app menu
		printMenu()

		input := inputs.GetUserInput("You choise: ")
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		switch num {
		case 1:
			// Create a new user object with firstname and created date
			u = users.NewUser()
		case 2:
			// Add the lastname to the User{}
			u.AddLastName()
		case 3:
			// Add the birthday to the User{}
			u.AddBirthDay()
		case 4:
			u.NewPrintUserInfo()
		case 5:
			// Change the fistname to the User{}
			u.ChangeFirstName()
		case 6:
			// Change the lastname to the User{}
			u.ChangeLastName()
		case 7:
			// Exit the app
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("1. Create a new user")
	fmt.Println("2. Add last name")
	fmt.Println("3. Add birth day")
	fmt.Println("4. Print user information")
	fmt.Println("5. Change the first name")
	fmt.Println("6. Change last name")
	fmt.Println("7. Exit")
	fmt.Println()
}
