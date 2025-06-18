package main

import (
	"context"
	"encoding/json"
	config "halo/cmd/adhoc/infrastructure"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Coordinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Location struct {
	Id      float32     `json:"id"`
	Name    string      `json:"name"`
	State   string      `json:"state"`
	Country string      `json:"country"`
	Coord   Coordinates `json:"coord"`
}

//lint:ignore U1000 this is a placeholder main function
func main() {
	dbpool, err := pgxpool.New(context.Background(), config.POSTGRES_URL)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	file, err := os.Open("./city.list.json")
	if err != nil {
		log.Printf("Unable to read file")
	}

	bytevalue, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Unable to parse file")
	}

	var locations []Location

	if err = json.Unmarshal(bytevalue, &locations); err != nil {
		log.Printf("Unable to unmarshal file %s", err)
	}

	createTables(dbpool)

	for _, item := range locations {
		insertLocation(dbpool, item)
		log.Printf("Item: %d, %s, %v", int(item.Id), item.Name, item.Coord.Lat)
	}

	log.Println("Completed")
}

func createTables(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS location (id INTEGER, name VARCHAR, state VARCHAR, country VARCHAR,latitude decimal, longitude decimal)")
	if err != nil {
		log.Printf("Error creating the tables %s", err)
	}
}

func insertLocation(db *pgxpool.Pool, item Location) {
	_, err := db.Exec(context.Background(), "INSERT INTO location (id, name, state, country, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6)", int(item.Id), item.Name, item.State, item.Country, item.Coord.Lat, item.Coord.Lon)
	if err != nil {
		log.Printf("Error inserting %s", item.Name)
	}
	log.Printf("Inserting")
}
