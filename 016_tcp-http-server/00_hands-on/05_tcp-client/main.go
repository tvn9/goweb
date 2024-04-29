package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello World from TCP client."))
}
