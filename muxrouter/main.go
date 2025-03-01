package main

import (
	"fmt"
	"log"
	"net/http"
)

// server struct holds server port address
type appConfig struct {
	addr string
}

func (a *appConfig) viewUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of Users..."))
}

func (a *appConfig) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new user"))
}

// main start the program
func main() {
	appCfg := &appConfig{
		addr: ":8080",
	}

	// mux router
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    appCfg.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", appCfg.viewUsersHandler)
	mux.HandleFunc("POST /createusers", appCfg.createUsersHandler)

	fmt.Println("test server from browser address - http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
