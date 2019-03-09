package database

import (
	"database/sql"
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
