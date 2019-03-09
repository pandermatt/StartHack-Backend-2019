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

func main() {
	addCars()
	router := mux.NewRouter()
	router.HandleFunc("/login", checkLogin).Methods("POST")
	router.HandleFunc("/rent/{id}", rentCar).Methods("POST")
	router.HandleFunc("/cars", getCars).Methods("GET")
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
	var p rental.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.PW != "test" {
		log.Fatalf("error: %e or wrong password", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("user %s attempted login", p.Name)
	w.WriteHeader(http.StatusOK)
}

// rent a car
func rentCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, v := range cars {
		if params["id"] == v.ID {
			if cars[i].Rented == true {
				cars[i].Rented = false
			} else {
				cars[i].Rented = true
			}
			log.Printf("car %s is now %t", cars[i].Name, cars[i].Rented)
		}
	}
	w.WriteHeader(http.StatusOK)
}

// get all the cars
func getCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}
