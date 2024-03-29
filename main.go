package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pandermatt/StartHackBackend/pkg/rental"
)

const (
	port = ":8000"
)

var cars []rental.Car
var person rental.Person
var reduction rental.Reduction

func main() {
	addCars()
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

// add the standard cars
func addCars() {
	cars = []rental.Car{}
	cars = append(cars, rental.Car{ID: "1", Name: "Volvo Typ 1", Rented: false, Image: "https://www.telegraph.co.uk/cars/images/2017/02/15/Volvo-S90-main_trans_NvBQzQNjv4BqLXKfuYoUkiu2TOJRKe-bQKhpKWSvo7bYwCFSVLx1AKs.jpg"})
	cars = append(cars, rental.Car{ID: "2", Name: "Volvo Typ 2", Rented: false, Image: "https://media.wired.com/photos/59e65901a00183307dad41f3/master/w_2400,c_limit/VolvoPolestarTA.jpg"})
	cars = append(cars, rental.Car{ID: "3", Name: "Volvo Typ 3", Rented: false, Image: "https://pictures.dealer.com/s/sandbergnorthwestvolvovcna/0217/c9c8d64fa9f413ea7a8d514bb5749ca9x.jpg?impolicy=resize&w=1024"})
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
	tmp := rental.Reduction{Clean: 0, Fueled: 0}
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		log.Fatal(err)
	}
	if &tmp != nil && &tmp.Clean != nil {
		reduction.Clean += tmp.Clean
	}
	if &tmp != nil && &tmp.Fueled != nil {
		reduction.Fueled += tmp.Fueled
	}
	reduction.Clean += tmp.Clean
	reduction.Fueled += tmp.Fueled
	json.NewEncoder(w).Encode(reduction)
}

// getReduction returns the reduction
func getReduction(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reduction)
}
