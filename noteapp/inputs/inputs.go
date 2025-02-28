package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// GetUserInput gets the keyboard input string
func GetUserInputReadString(s string) (string, error) {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return line, err
	} else {
		return line, nil
	}
}

func GetNoteTitle(s string) (string, error) {
	var text string
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text = scanner.Text()
	return text, nil
}

func GetNoteContent(s string) (string, error) {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	text := scanner.Text()
	return text, nil
}

// InputErrorCheck verify for valid user input
func InputErrorCheck(input string) error {
	if len(input) < 3 {
		return errors.New("error: input length must be atleast 3 letters")
	}
	return nil
}
