package repository

import (
	"context"
	"halo/internal/app"
	"log"
	"strconv"
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
	var namesFormatted []interface{}

	var whereClause string

	for i, name := range namesUnformatted {

		name = name + "%"
		namesFormatted = append(namesFormatted, name)
		if i != 0 {
			whereClause = whereClause + " or"
		}
		whereClause = whereClause + " LOWER(name) LIKE LOWER($" + strconv.Itoa(i+1) + ") or LOWER(state) LIKE LOWER($" + strconv.Itoa(i+1) + ")"
	}

	sqlQuery := "SELECT * FROM location WHERE" + whereClause
	log.Printf("%s", sqlQuery)
	rows, err := db.Query(context.Background(), sqlQuery, namesFormatted...)
	if err != nil {
		log.Printf("There was an error in the query %s", err)
	}

	for rows.Next() {
		location := &app.Location{}

		if err := rows.Scan(&location.Id, &location.Name, &location.State, &location.Latitude, &location.Longitude); err != nil {
			log.Printf("scan error for row")
		}

		locations = append(locations, location)
	}

	return locations
}
