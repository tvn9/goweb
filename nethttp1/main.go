// Understanding net/http package
// exmaple 1
package main

import (
	"fmt"
	"io"
	"net/http"
)

type test struct{}

func (a test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Writing from fmt.Fprintln!")
	w.Write([]byte(`Writing string from w.Write([]byte("string"))`))

	// using w.Header().Set()
	// Setting Trailer
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Set("Trailer", "AtEnd3")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

	w.Header().Set("AtEnd1", "value 1")
	io.WriteString(w, "\nThis HTTP response has both headers before this text and trailers at the end.\n")
	w.Header().Set("AtEnd2", "value 2")
	w.Header().Set("AtEnd3", "value 3") // These will appear as trailer.
}

func main() {
	a := test{}

	fmt.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", a)
}
