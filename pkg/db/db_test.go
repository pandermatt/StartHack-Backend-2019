package database

import (
	"testing"

	"github.com/pandermatt/StartHackBackend/pkg/rental"
)

func TestDb(t *testing.T) {
	db := OpenDB()
	tx, err := db.Begin()
	if err != nil {
		t.Log(err)
	}
	AddCars(db)
	tx.Commit()
	defer db.Close()
}

func TestDbContent(t *testing.T) {
	db := OpenDB()
	rows, err := db.Query("select id, name, image, rented from cars")
	if err != nil {
		t.Log(err)
	}
	AddCars(db)
	defer rows.Close()
	for rows.Next() {
		var car rental.Car
		err = rows.Scan(&car.ID, &car.Name, &car.Image, &car.Rented)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(car)
	}
}
