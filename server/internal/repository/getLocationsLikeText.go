package repository

import (
	"context"
	"halo/internal/app"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationsLikeText(db *pgxpool.Pool, loc string) []*app.Location {
	var locations []*app.Location

	namesUnformatted := strings.Split(loc, ",")
	if len(namesUnformatted) > 2 {
		log.Printf("Too many commas")
		return locations
	}

	var formatted []any = make([]any, len(namesUnformatted))
	i := 0
	for range len(namesUnformatted) {
		item := "%" + namesUnformatted[i] + "%"
		formatted = append(formatted, item)
	}

	var whereClause string

	whereClause = whereClause + " LOWER(city) LIKE ANY (array[$1"
	if len(namesUnformatted) > 1 {
		whereClause = whereClause + ", $2])"
	} else {
		whereClause = whereClause + "])"
	}
	whereClause = whereClause + " and (LOWER(state_name) LIKE ANY (array[$1"
	if len(namesUnformatted) > 1 {
		whereClause = whereClause + ", $2])"
	} else {
		whereClause = whereClause + "])"
	}
	whereClause = whereClause + " or LOWER(city) LIKE ANY (array[$1"
	if len(namesUnformatted) > 1 {
		whereClause = whereClause + ", $2]))"
	} else {
		whereClause = whereClause + "]))"
	}

	sqlQuery := "SELECT DISTINCT id, city as name, state_id as state, lat as latitude, lng as longitude, ranking FROM location WHERE" + whereClause + " ORDER BY ranking LIMIT 50"

	rows, err := db.Query(context.Background(), sqlQuery, formatted[len(namesUnformatted):]...)

	if err != nil {
		log.Printf("There was an error in the query %s", err)
	}

	for rows.Next() {
		location := &app.Location{}

		if err := rows.Scan(&location.Id, &location.Name, &location.State, &location.Latitude, &location.Longitude, &location.Ranking); err != nil {
			log.Printf("scan error for row %s", err)
		}

		locations = append(locations, location)
	}

	return locations
}
