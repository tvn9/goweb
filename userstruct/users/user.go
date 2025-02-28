package users

import (
	"fmt"
	"time"
	"userstruct/inputs"
)

// Declare a veriable with struct data type
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt string
}

// NewUser function add the firstname to the u (user object)
func NewUser() User {
	u := User{}
	fname := inputs.GetUserInput("First name: ")
	if err := inputs.InputErrorCheck(fname); err != nil {
		fmt.Println(err)
		return u
	}
	u.firstName = fname
	u.createdAt = time.Now().Format(time.RFC3339)
	return u
}

// addLastName function assigns the lastname to the existing user
func (u *User) AddLastName() {
	lname := inputs.GetUserInput("Last name: ")
	if err := inputs.InputErrorCheck(lname); err != nil {
		fmt.Println(err)
	} else {
		u.lastName = lname
	}
}

// addBirthDay function assigns the birthday to the existing user{} bDate field
func (u *User) AddBirthDay() {
	bday := inputs.GetUserInput("Birthday: ")
	if err := inputs.InputErrorCheck(bday); err != nil {
		fmt.Println(err)
	} else {
		u.birthDate = bday
	}
}

// changeFirstName method change the firstname (fName) from the existing user object
func (u *User) ChangeFirstName() {
	fname := inputs.GetUserInput("Change firstname to: ")
	if err := inputs.InputErrorCheck(fname); err != nil {
		fmt.Println(err)
	} else {
		u.firstName = fname
	}
}

// changeLastName method change the firstname (fName) from the existing user object
func (u *User) ChangeLastName() {
	lname := inputs.GetUserInput("Change last name to: ")
	if err := inputs.InputErrorCheck(lname); err != nil {
		fmt.Println(err)
	} else {
		u.lastName = lname
	}
}

// NewPrintUserInfo method print the user{} fields to the screen
func (u User) NewPrintUserInfo() {
	fmt.Println()
	fmt.Println("First name:", u.firstName)
	fmt.Println("Last name:", u.lastName)
	fmt.Println("Birth day:", u.birthDate)
	fmt.Println("Recorded time:", u.createdAt)
	fmt.Println()
}
