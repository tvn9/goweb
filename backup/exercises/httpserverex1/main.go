// exercise - create the http server that will response the URL of the request
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write respond
	respond(conn)
}

// request receives the http request, scan the content of the request
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("Requested URL:", u)
			fmt.Println("***METHOD***", m)
		}
		if ln == "" {
			// header are done
			break
		}
		i++
	}
}

// respond returns the http header information and the contest from the
// request URL
func respond(conn net.Conn) {
	body := `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	<h1>Exercise - http server that respond the requested URL</h1>
</body>
</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handle(conn)
	}
}
