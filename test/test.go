package main

import (
	"bufio"
	"fmt"
	"os"
)

type car struct {
	model string
}

func main() {
	sliceTest()
	// arrayTest()

	// str, err := readString("Enter some text: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(str)

	// str, err = scannerString("Enter some text: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(str)
}

func printType(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Println("Int:", v)
	case float32:
		fmt.Printf("Float32: %T, %v\n", v, v)
	case float64:
		fmt.Printf("Float64: %T, %v\n", v, v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Type not supported")
	}
}

func printValidedType(v interface{}) {
	val, ok := v.(person)
	if ok {
		fmt.Printf("%T, %v\n", val, val)
	} else {
		fmt.Println("wrong type")
	}
}

func readString(s string) (string, error) {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return line, err
	} else {
		return line, nil
	}
}

func scannerString(s string) (string, error) {
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
