package repository

import (
	"context"
	"halo/internal/app"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationsWithinCoords(db *pgxpool.Pool, maxLat float32, minLat float32, maxLong float32, minLong float32) []*app.Location {
	var locations []*app.Location

	sqlQuery := "SELECT * FROM location WHERE latitude BETWEEN $1 and $2 and longitude BETWEEN $3 and $4"

	rows, err := db.Query(context.Background(), sqlQuery, minLat, maxLat, minLong, maxLong)

	if err != nil {
		log.Printf("Error retrieving locations within coordinates")
	}

	for rows.Next() {
		location := &app.Location{}

		if err := rows.Scan(&location.Id, &location.Name, &location.State, &location.Country, &location.Coord.Lat, &location.Coord.Lon); err != nil {
			log.Printf("scan error for row")
		}

		locations = append(locations, location)
	}

	return locations
}
