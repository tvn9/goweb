package main

import (
	"fmt"
	"net/http"
)

type Name string

const name = `
<!doctype html>
<html>
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width" />
        <title>Document</title>
    </head>
    <body>
        <h1>My name is Thanh Nguyen</h1>
    </body>
</html>`

func (n Name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, name)
}

func main() {
	var n Name
	_ = http.ListenAndServe(":8080", n)
}
