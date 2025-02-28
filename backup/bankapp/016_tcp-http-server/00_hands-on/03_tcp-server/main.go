package main

import (
	"log"
	"net"
)

func main() {
	// start the tcp server
	li, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		// accept a connection
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// process connection
		func(c net.Conn) {
			// setup []byte reading buffer
			buf := make([]byte, 1024)
			// read data from connection and store in buffer
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			// print convert byte from buf to string and print to stdout
			log.Print(string(buf))

			// convert string to []byte and print to connection
			conn.Write([]byte("Hello from TCP server"))
		}(conn)
	}
}
