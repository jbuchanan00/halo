package repository

import (
	"context"
	"halo/internal/app"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationsLikeText(db *pgxpool.Pool, loc string) []*app.Location {
	var locations []*app.Location

	sqlQuery := "SELECT * FROM location WHERE LOWER(name) LIKE $1 or LOWER(state) LIKE $1"

	rows, err := db.Query(context.Background(), sqlQuery, loc+"%")
	if err != nil {
		log.Printf("There was an error in the query %s", err)
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
