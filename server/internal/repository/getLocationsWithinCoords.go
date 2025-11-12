package repository

import (
	"context"
	"halo/internal/app"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationsWithinCoords(db *pgxpool.Pool, maxLat float32, minLat float32, maxLong float32, minLong float32) []*app.Location {
	var locations []*app.Location

	sqlQuery := "SELECT id, city_name as name, state_id as state, lat as latitude, lng as longitude, ranking FROM location WHERE (lat BETWEEN $1 and $2) and (lng BETWEEN $3 and $4) ORDER BY ranking"

	rows, err := db.Query(context.Background(), sqlQuery, float64(minLat), float64(maxLat), float64(minLong), float64(maxLong))

	if err != nil {
		log.Printf("Error retrieving locations within coordinates")
	}

	for rows.Next() {
		location := &app.Location{}

		if err := rows.Scan(&location.Id, &location.Name, &location.State, &location.Latitude, &location.Longitude, &location.Ranking); err != nil {
			log.Printf("scan error for row")
		}

		locations = append(locations, location)
	}

	return locations
}
