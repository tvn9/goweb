package main

import (
	"fmt"
	"net"
)

func main() {
	l := "localhost"
	p := "8080"
	hostport := net.JoinHostPort(l, p)
	fmt.Println(hostport)

	addreses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, a := range addreses {
		fmt.Printf("%d - %s\n", i, a)

	}
}
