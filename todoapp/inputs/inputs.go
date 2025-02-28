package inputs

import (
	"bufio"
	"fmt"
	"os"
	"todoapp/errorcheck"
)

// GetUserInput gets the keyboard input string
func GetUserInputReadString(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := errorcheck.InputErrorCheck(line); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return line
}

func ScanUserInput(s string) string {
	fmt.Print(s)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	text := scanner.Text()
	if err := errorcheck.InputErrorCheck(text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return text
}
