package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Please enter a name")
		return
	}
	name := strings.Join(args, " ")
	str := fmt.Sprint(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
    </head>
<body>
    <h1>` + name + `</h1>
</body>
</html>
`)
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("fails to create file - ", err)
	}

	if _, err := file.WriteString(str); err != nil {
		log.Fatal("fails to write to file - ", err)
	}
}
