package main

import (
	"io"
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
		go func(c net.Conn) {
			bytes := []byte{}
			for {
				// setup []byte reading buffer
				buf := make([]byte, 32)
				// read data from connection and store in buffer
				_, err := c.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					} else {
						log.Fatal(err)
					}
				}
				// print convert byte from buf to string and print to stdout
				log.Print(string(buf))
				// adding data from each connection to byte slice
				bytes = append(bytes, buf...)
				// convert string to []byte and print to connection
				conn.Write([]byte("Hello from TCP server\r\n"))
				// writing bytes []byte to connection
				conn.Write(bytes)
			}
		}(conn)
	}
}
