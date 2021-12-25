// Hello world of Golang!
package main

import (
	"fmt"
	"net/http"
)

// handler send the target string to http writer
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world of Golang!, %s\n", r.URL.Path)
}

func main() {

	// call handler()
	http.HandleFunc("/", handler)

	// start the built-in web server
	http.ListenAndServe(":8080", nil)

}
