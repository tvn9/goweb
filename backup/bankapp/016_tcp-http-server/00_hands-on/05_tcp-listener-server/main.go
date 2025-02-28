package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	li, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal()
			}
			log.Print(string(buf))
			conn.Write([]byte("Hello from TCP server"))
			c.Close()
		}(conn)
	}
}
