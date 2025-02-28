package errorcheck

import (
	"errors"
)

// InputErrorCheck verify for valid user input
func InputErrorCheck(input string) error {
	if input == "" || len(input) < 3 {
		return errors.New("error: emplty text is not supported, input must be greater than 3 letters")
	} else {
		return nil
	}
}
