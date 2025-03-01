package main

import "net/http"

type server struct {
	addr string
}

func main() {
	s := &server{
		addr: ":8080",
	}
	_ = http.ListenAndServe(s.addr, s)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := []byte("Hello From the server!")
	_, _ = w.Write(str)
}
