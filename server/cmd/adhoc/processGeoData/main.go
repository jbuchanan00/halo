package processGeoData

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

type NewLocation struct {
	State     string  `json:"state"`
	City      string  `json:"city"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

//lint:ignore U1000 this is a placeholder main function
func ProcessGeoData() {
	dbpool, err := pgxpool.New(context.Background(), config.GetPostgresUrl())
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	file, err := os.Open("./cityStateToLatLong.json")
	if err != nil {
		log.Printf("Unable to read file")
	}

	bytevalue, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Unable to parse file")
	}

	var locations []NewLocation

	if err = json.Unmarshal(bytevalue, &locations); err != nil {
		log.Printf("Unable to unmarshal file %s", err)
	}

	createTables(dbpool)
	id := 0
	for _, item := range locations {
		id++
		insertLocation(dbpool, item, id)
	}

	log.Println("Completed")
}

func createTables(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS location (id INTEGER, name VARCHAR, state VARCHAR, latitude decimal, longitude decimal)")
	if err != nil {
		log.Printf("Error creating the tables %s", err)
	}
}

func insertLocation(db *pgxpool.Pool, item NewLocation, id int) {
	_, err := db.Exec(context.Background(), "INSERT INTO location (id, name, state, latitude, longitude) VALUES ($1, $2, $3, $4, $5)", id, item.City, item.State, item.Latitude, item.Longitude)
	if err != nil {
		log.Printf("Error inserting %s", item.City)
	}
	log.Printf("Inserting")
}
