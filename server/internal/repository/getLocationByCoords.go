package repository

import (
	"context"
	"halo/internal/app"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationByCoords(db *pgxpool.Pool, coords *app.Coordinates) *app.Location {
	location := &app.Location{}
	sqlQuery := "SELECT id, city as name, state_id as state, lat as latitude, lng as longitude, ranking FROM location WHERE lat between $1 and $2 and lng between $3 and $4"

	params := []any{
		float64(coords.Latitude - 0.001),
		float64(coords.Latitude + 0.001),
		float64(coords.Longitude - 0.001),
		float64(coords.Longitude + 0.001)}

	row := db.QueryRow(context.Background(), sqlQuery, params...)

	if err := row.Scan(&location.Id, &location.Name, &location.State, &location.Latitude, &location.Longitude, &location.Ranking); err != nil {
		log.Printf("Scan error for row %s", err)
	}

	return location
}
