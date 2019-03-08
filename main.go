package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":8000"
)

type person struct {
	PW   string `json:"pw"`
	Name string `json:"name"`
}

type rental struct {
	Duration string `json:"duration"`
	Name     string `json:"name"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", checkLogin).Methods("GET")
	router.HandleFunc("/rent", rentCar).Methods("POST")
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	var p person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.PW != "test" {
		log.Fatalf("error: %e or wrong password", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("user %s now logged in", p.Name)
	w.WriteHeader(http.StatusOK)
}

func rentCar(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world")
}
