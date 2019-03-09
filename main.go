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
	var person rental.Person
	l := rental.Login{Correct: true}
	_ = json.NewDecoder(r.Body).Decode(&person)
	if !database.UserExists(person.Name, db) {
		l.Correct = false
		// insert new user, the password will be set to the one
		// he entered
		database.InsertNewUser(person.Name, person.PW, db)
		log.Print("inserted new user")
	} else {
		if !database.CheckUserPw(person.Name, person.PW, db) {
			l.Correct = false
		}
		log.Print("checked existing user against his password")
	}
	json.NewEncoder(w).Encode(l)
}

// rent a car
func rentCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	index := 0
	cars := database.GetCars(db)
	for i, v := range cars {
		if params["id"] == v.ID {
			index = i
		}
	}
	database.UpdateCar(cars[index].ID, !cars[index].Rented, db)
	json.NewEncoder(w).Encode(cars[index])
}

// get all the cars
func getCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.GetCars(db))
}

// subs adds a subscription
func subs(w http.ResponseWriter, r *http.Request) {
	var person rental.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	database.UpdateSubs(person.Name, person.Subscription, db)
	json.NewEncoder(w).Encode(person)
}

// reduce remarks the points a user has accumulated
func reduce(w http.ResponseWriter, r *http.Request) {
	var reduction rental.Reduction
	err := json.NewDecoder(r.Body).Decode(&reduction)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reduction)
	json.NewEncoder(w).Encode(reduction)
}

// getReduction returns the reduction
func getReduction(w http.ResponseWriter, r *http.Request) {
	var reduction rental.Reduction
	json.NewEncoder(w).Encode(reduction)
}
