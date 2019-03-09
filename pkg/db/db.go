package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver for sqlite
	_ "github.com/mattn/go-sqlite3"
	"github.com/pandermatt/StartHackBackend/pkg/rental"
)

const (
	tblUser      = "create table if not exists persons (name text not null primary key, pw text not null, subscription text not null)"
	tblCar       = "create table if not exists cars (id text primary key, image text not null, name text not null, rented boolean not null)"
	tblReduction = "create table if not exists reductions (name text primary key, clean integer not null, fueled integer not null, foreign key (name) references person(name))"
)

// OpenDB creates a db and opens it
func OpenDB() *sql.DB {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(tblUser)
	_, err = db.Exec(tblCar)
	_, err = db.Exec(tblReduction)

	if err != nil {
		defer db.Close()
		panic(err)
	}

	return db
}

// AddCars adds the cars to the database
func AddCars(db *sql.DB) {
	cars := []rental.Car{}
	cars = append(cars, rental.Car{ID: "1", Name: "Volvo Typ 1", Rented: false, Image: "https://www.telegraph.co.uk/cars/images/2017/02/15/Volvo-S90-main_trans_NvBQzQNjv4BqLXKfuYoUkiu2TOJRKe-bQKhpKWSvo7bYwCFSVLx1AKs.jpg"})
	cars = append(cars, rental.Car{ID: "2", Name: "Volvo Typ 2", Rented: false, Image: "https://media.wired.com/photos/59e65901a00183307dad41f3/master/w_2400,c_limit/VolvoPolestarTA.jpg"})
	cars = append(cars, rental.Car{ID: "3", Name: "Volvo Typ 3", Rented: false, Image: "https://pictures.dealer.com/s/sandbergnorthwestvolvovcna/0217/c9c8d64fa9f413ea7a8d514bb5749ca9x.jpg?impolicy=resize&w=1024"})
	stm, err := db.Prepare("insert into cars (id, image, name, rented) values (?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	for _, v := range cars {
		_, err = stm.Exec(v.ID, v.Image, v.Name, v.Rented)
		if err != nil {
			log.Print(err)
		}
	}
}

// GetCars returns all the cars in the database
func GetCars(db *sql.DB) []rental.Car {
	cars := []rental.Car{}
	rows, err := db.Query("select id, name, image, rented from cars")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var car rental.Car
		err = rows.Scan(&car.ID, &car.Name, &car.Image, &car.Rented)
		if err != nil {
			log.Fatal(err)
		}
		cars = append(cars, car)
	}
	return cars
}

// UserExists checks if the user already exists
func UserExists(u string, db *sql.DB) bool {
	rows, err := db.Query(fmt.Sprintf("select name, pw, subscription from users where name = %s", u))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	people := []rental.Person{}
	for rows.Next() {
		var person rental.Person
		err = rows.Scan(&person.Name, &person.PW, &person.Subscription)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}
	return len(people) >= 1
}

// CheckUserPw checks if the password matches
// the one in the database
func CheckUserPw(u, p string, db *sql.DB) bool {
	rows, err := db.Query(fmt.Sprintf("select name, pw, subscription from users where name = %s", u))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	people := []rental.Person{}
	for rows.Next() {
		var person rental.Person
		err = rows.Scan(&person.Name, &person.PW, &person.Subscription)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}
	return people[0].PW == p
}

// InsertNewUser inserts a new user with the given password
func InsertNewUser(u, p string, db *sql.DB) {
	_, err := db.Query(fmt.Sprintf("insert into persons (name, pw, subscription) values (%s, %s, %s)", u, p, ""))
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateCar updates the state of the car
func UpdateCar(id string, rented bool, db *sql.DB) {
	_, err := db.Query(fmt.Sprintf("update cars set rental = %t where id = %s", rented, id))
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateSubs updates the state of the car
func UpdateSubs(name, subs string, db *sql.DB) {
	_, err := db.Query(fmt.Sprintf("update persons set subscription = %s where name = %s", subs, name))
	if err != nil {
		log.Fatal(err)
	}
}
