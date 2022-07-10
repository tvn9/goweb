// Understanding net/http ServerMux package
// Routing example 1

package main

import "net/http"

type someThing struct{}

func (s someThing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home":
		fallthrough
	case "/":
		w.Write([]byte("<h1>Welcome to homepage!</h1>"))
	case "/about":
		w.Write([]byte("<h1>You are now on about page.</h1>"))
	}
}

func main() {
	var s someThing

	http.ListenAndServe(":8080", s)
}
