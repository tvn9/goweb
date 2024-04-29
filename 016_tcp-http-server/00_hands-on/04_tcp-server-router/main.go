package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var (
	url  = "/"
	form = `<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>`
)

func printData(conn net.Conn, url, applyForm string) {
	body := `<!DOCTYPE html>
    <html lang="en">
        <head><meta charset="UTF-8">
            <title></title>
        </head>
        <body>
            <strong>` + url + `</strong><br>
            <a href="/">index</a><br>
            <a href="/about">about</a><br>
            <a href="/contact">contact</a><br>
            <a href="/apply">apply</a><br>` + applyForm +
		`</body>
    </html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	url = "INDEX"
	applyForm := ""
	printData(conn, url, applyForm)
}

func about(conn net.Conn) {
	url = "ABOUT"
	applyForm := ""
	printData(conn, url, applyForm)
}

func contact(conn net.Conn) {
	url = "CONTACT"
	applyForm := ""
	printData(conn, url, applyForm)
}

func apply(conn net.Conn) {
	url = "APPLY"
	applyForm := form
	printData(conn, url, applyForm)
}

func applyProcess(conn net.Conn) {
	url := "APPLY PROCESS"
	applyForm := ""
	printData(conn, url, applyForm)
}
