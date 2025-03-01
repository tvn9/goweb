package main

// Writing a basic router to understand basic http routing

import (
	"fmt"
	"net/http"
)

// server struct holds server port address
type appConfig struct {
	addr string
}

// ServerHTTP is the function from the (handler interface)
func (s *appConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index page\n"))
		case "/contact":
			w.Write([]byte("Welcome to contact page!\n"))
		case "/about":
			w.Write([]byte("Welcome to about us page!\n"))
		case "/users":
			w.Write([]byte("View all users\n"))
		default:
			w.Write([]byte("404 page not found\n"))
		}
	}
}

// main start the program
func main() {
	appCfg := &appConfig{
		addr: ":8080",
	}

	srv := &http.Server{
		Addr:    appCfg.addr,
		Handler: appCfg,
	}

	fmt.Println("test server from browser address - http://localhost:8080")
	srv.ListenAndServe()
}
