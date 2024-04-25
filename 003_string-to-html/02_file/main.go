package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := "Thanh Nguyen"
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

	file, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("fails to open or create file, error:", err)
	}
	defer file.Close()

	_, err = file.WriteString(str)
	if err != nil {
		log.Fatal("fails to write to file, error:", err)
	}
}
