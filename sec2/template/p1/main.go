// Creating HTML template with string text
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
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
</body>
</html>
`

	fmt.Println(tpl)

	// Create a file to save as index.HTML
	index, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer index.Close()

	io.Copy(index, strings.NewReader(tpl))
}
