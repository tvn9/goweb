// Servve Content - Using http.ServeFile

package main

import "net/http"

func pix(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}

func main() {
	http.HandleFunc("/", pix)
	http.ListenAndServe(":8080", nil)
}
