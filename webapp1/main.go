package main

import (
	"log"
	"net/http"
)

const portNum = ":8080"

// main
func main() {
	// Index
	http.HandleFunc("/", index)

	// About
	http.HandleFunc("/about", about)

	// Start the web server on port :8080
	log.Println("Started server on port:", portNum)
	http.ListenAndServe(":8080", nil)

}
