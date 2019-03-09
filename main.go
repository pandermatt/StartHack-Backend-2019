package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pandermatt/StartHackBackend/pkg/db"
	"github.com/pandermatt/StartHackBackend/pkg/rental"
)

const (
	port = ":8000"
)

var cars []rental.Car
var person rental.Person
var reduction rental.Reduction
var db *sql.DB

func main() {
	db = database.OpenDB()
	defer db.Close()
	database.AddCars(db)
	router := mux.NewRouter()
	router.HandleFunc("/login", checkLogin).Methods("POST")
	router.HandleFunc("/rent/{id}", rentCar).Methods("POST")
	router.HandleFunc("/cars", getCars).Methods("GET")
	router.HandleFunc("/reduction", reduce).Methods("POST")
	router.HandleFunc("/reduction", getReduction).Methods("GET")
	router.HandleFunc("/subscription", subs).Methods("POST")
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// check the login and return 200OK
func checkLogin(w http.ResponseWriter, r *http.Request) {
	l := rental.Login{Correct: true}
	_ = json.NewDecoder(r.Body).Decode(&person)
	if person.PW != "test" {
		l.Correct = false
		log.Print("login false")
	} else {
		log.Print("login true")
	}
	json.NewEncoder(w).Encode(l)
}

// rent a car
func rentCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	index := 0
	for i, v := range cars {
		if params["id"] == v.ID {
			index = i
			if cars[i].Rented == true {
				cars[i].Rented = false
			} else {
				cars[i].Rented = true
			}
			log.Printf("car %s is now %t", cars[i].Name, cars[i].Rented)
		}
	}
	json.NewEncoder(w).Encode(cars[index])
}

// get all the cars
func getCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}

// subs adds a subscription
func subs(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(person.Subscription)
	json.NewEncoder(w).Encode(person)
}

// reduce remarks the points a user has accumulated
func reduce(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&reduction)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reduction)
	json.NewEncoder(w).Encode(reduction)
}

// getReduction returns the reduction
func getReduction(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reduction)
}
