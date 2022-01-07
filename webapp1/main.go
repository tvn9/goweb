package main

import (
	"log"
	"net/http"
)

const portNum = ":8080"

// main
func main() {
	// Handlefunc
	http.HandleFunc("/", Index)
	http.HandleFunc("/about", About)

	// Start the web server on port :8080
	log.Println("Starting server on port", portNum)
	http.ListenAndServe(portNum, nil)

}
