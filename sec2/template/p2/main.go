// Creating HTML template with string text and save to a file
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// validate cli-argument
	if len(args) != 1 {
		fmt.Println("Please enter and file name.")
		return
	}

	title := "This is the home page!"
	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
</head>
<body>
	<h1>` + title + `</h1>
	<p>This html template will be saved to a file name input from command-line argument</p>
</body>
</html>
`

	// Println the template text to terminal
	fmt.Println(tpl)

	fileName := strings.ToLower(args[0])
	// Create a file to save as file name you have entered.html
	fn, err := os.Create(fileName + ".html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer fn.Close()

	nb, err := io.Copy(fn, strings.NewReader(tpl))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of bytes copied:", nb)
}
