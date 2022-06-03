package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	fmt.Println(uuid)

	http.HandleFunc("/", index)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		uuid := uuid.New()
		fmt.Println(uuid)
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
