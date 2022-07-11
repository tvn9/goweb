// # Create a basic server using TCP.
// The server should use net.Listen to listen on port 8080.
// Remember to close the listener using defer.
// Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.
// Now write a response back on the connection.
// Use io.WriteString to write the response: I see you connected.
// Remember to close the connection.
// Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Setup TCP connection
	log.Println("Starting TCP Server on Port: 8080")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// setup a connection scanner
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println("Testing", ln)
			if ln == "" {
				log.Println("No more data!")
				break
			}
			serve(conn)
		}
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("End of data")
			break
		}
	}
	io.WriteString(c, "Write data to connection.")
}
