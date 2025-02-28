package inputs

import (
	"errors"
	"fmt"
)

// GetUserInput gets the keyboard input string
func GetUserInput(s string) string {
	var value string
	fmt.Print(s)
	fmt.Scanln(&value)
	return value
}

// InputErrorCheck verify for valid user input
func InputErrorCheck(input string) error {
	if len(input) < 3 {
		return errors.New("Error: input length must be atleast 3 letters.")
	}
	return nil
}
