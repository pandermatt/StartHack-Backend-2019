package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":8000"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", checkLogin).Methods("GET")
	router.HandleFunc("/rent", rentCar).Methods("POST")
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func rentCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
