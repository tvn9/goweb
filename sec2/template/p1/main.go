// Creating HTML template with string text
package main

import "fmt"

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
}
