package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello, World!")
	w.Write(hello)
}
func main() {
	s := &server{
		addr: ":8080",
	}
	fmt.Println("test server from browser address - http://localhost:8080")
	log.Fatal(http.ListenAndServe(s.addr, s))
}
